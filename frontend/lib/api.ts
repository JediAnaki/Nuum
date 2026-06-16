import axios from 'axios';
import type { AuthResponse, Video, VideoListResponse } from '@/types';

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api/v1';

export const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Add auth token to requests
api.interceptors.request.use((config) => {
  const token = localStorage.getItem('access_token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

// Auth API
export const authApi = {
  register: async (data: { username: string; email: string; password: string; full_name?: string }) => {
    const response = await api.post<AuthResponse>('/auth/register', data);
    return response.data;
  },

  login: async (data: { email: string; password: string }) => {
    const response = await api.post<AuthResponse>('/auth/login', data);
    return response.data;
  },

  getMe: async () => {
    const response = await api.get('/auth/me');
    return response.data;
  },
};

// Video API
export const videoApi = {
  getVideos: async (page = 1, pageSize = 20, userId?: number) => {
    const params = new URLSearchParams({
      page: page.toString(),
      page_size: pageSize.toString(),
    });

    if (userId) {
      params.append('user_id', userId.toString());
    }

    const response = await api.get<VideoListResponse>(`/videos?${params}`);
    return response.data;
  },

  getVideo: async (id: number) => {
    const response = await api.get<Video>(`/videos/${id}`);
    return response.data;
  },

  uploadVideo: async (file: File, title: string, description?: string) => {
    const formData = new FormData();
    formData.append('video', file);
    formData.append('title', title);
    if (description) {
      formData.append('description', description);
    }

    const response = await api.post<Video>('/videos', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });
    return response.data;
  },

  deleteVideo: async (id: number) => {
    await api.delete(`/videos/${id}`);
  },
};
