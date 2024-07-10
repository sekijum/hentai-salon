import axios from 'axios';

export const useActions = () => {
  const { $api } = useNuxtApp();

  async function fetchListPresignedUrl(objectNames: string[]): Promise<string[]> {
    const response = await $api.post<{ urls: string[] }>('/files/urls-for-upload', { objectNames });
    return response.data.urls;
  }

  async function uploadFileToS3WithPresignedUrl(presignedUrl: string, file: File): Promise<string> {
    const response = await axios.put(presignedUrl, file, {
      headers: {
        'Content-Type': file.type,
      },
    });
    console.log(response);
    if (response.status !== 200) {
      throw new Error('Failed to upload file to S3');
    }

    return presignedUrl.split('?')[0];
  }

  return { fetchListPresignedUrl, uploadFileToS3WithPresignedUrl };
};
