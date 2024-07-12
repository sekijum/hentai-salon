import axios from 'axios';
import type { IThreadCommentAttachment } from '~/types/thread-comment-attachment';

export const useActions = () => {
  const { $api } = useNuxtApp();

  async function fetchListPresignedUrl(objectNames: string[]): Promise<string[]> {
    const response = await $api.post<{ urls: string[] }>('/files/urls-for-upload', { objectNames });
    return response.data.urls;
  }

  async function uploadFilesToS3(presignedUrl: string, file: File): Promise<string> {
    const response = await axios.put(presignedUrl, file, {
      headers: {
        'Content-Type': file.type,
      },
    });
    if (response.status !== 200) {
      throw new Error('Failed to upload file to S3');
    }
    return presignedUrl.split('?')[0];
  }

  async function uploadFilesToImgur(files: File[]): Promise<IThreadCommentAttachment[]> {
    const config = useRuntimeConfig();

    const uploadedAttachments: IThreadCommentAttachment[] = [];
    const clientId = config.public.clientId;

    for (let i = 0; i < files.length; i++) {
      const file = files[i];
      const formData = new FormData();
      const type = file.type.startsWith('video') ? 'Video' : 'Image';

      let uploadUrl = 'https://api.imgur.com/3/image';
      if (type === 'Video') {
        formData.append('video', file);
        uploadUrl = 'https://api.imgur.com/3/upload';
      } else {
        formData.append('image', file);
      }

      try {
        const response = await axios.post(uploadUrl, formData, {
          headers: {
            Authorization: `Client-ID ${clientId}`,
          },
        });
        const data = response.data;
        if (data.success) {
          uploadedAttachments.push({
            url: data.data.link,
            displayOrder: i,
            type: type,
          });
        } else {
          console.error(`Failed to upload file: ${file.name}`, data);
        }
      } catch (error) {
        console.error(`Error uploading file: ${file.name}`, error);
      }
    }

    return uploadedAttachments;
  }

  async function uploadFilesToVimeo(files: File[]): Promise<IThreadCommentAttachment[]> {
    const config = useRuntimeConfig();
    const uploadedAttachments: IThreadCommentAttachment[] = [];
    const accessToken = config.public.vimeoAccessToken;

    for (let i = 0; i < files.length; i++) {
      const file = files[i];
      const formData = new FormData();
      formData.append('file_data', file);

      const type = file.type.startsWith('video') ? 'Video' : 'Image';

      try {
        const createResponse = await axios.post('https://api.vimeo.com/me/videos', null, {
          headers: {
            'Authorization': `Bearer ${accessToken}`,
            'Content-Type': 'application/json',
            'Accept': 'application/vnd.vimeo.*+json;version=3.4',
          },
        });

        const uploadLink = createResponse.data.upload.upload_link;
        await axios.put(uploadLink, file, {
          headers: {
            'Content-Type': file.type,
          },
        });

        const videoUrl = createResponse.data.link;

        uploadedAttachments.push({
          url: videoUrl,
          displayOrder: i,
          type: type,
        });
      } catch (error) {
        console.error('Vimeo upload error:', error);
      }
    }

    return uploadedAttachments;
  }

  return { fetchListPresignedUrl, uploadFilesToS3, uploadFilesToImgur, uploadFilesToVimeo };
};
