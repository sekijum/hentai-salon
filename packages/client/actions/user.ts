import api from '~/utils/api';

export async function getAuthenticatedUser() {
  try {
    const response = await api.get('/whoami');
    return response.data;
  } catch (error) {
    throw error;
  }
}
