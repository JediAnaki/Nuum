export interface User {
  id: number;
  username: string;
  email: string;
  full_name: string;
  avatar?: string;
  bio?: string;
  is_active: boolean;
  created_at: string;
  updated_at: string;
}

export interface Video {
  id: number;
  user_id: number;
  title: string;
  description: string;
  thumbnail?: string;
  duration: number;
  status: 'uploading' | 'processing' | 'ready' | 'failed';
  views_count: number;
  likes_count: number;
  width: number;
  height: number;
  codec: string;
  bitrate: number;
  created_at: string;
  updated_at: string;
  user?: User;
  processed_videos?: ProcessedVideo[];
}

export interface ProcessedVideo {
  id: number;
  video_id: number;
  quality: string;
  path: string;
  size: number;
  bitrate: number;
  created_at: string;
}

export interface Comment {
  id: number;
  video_id: number;
  user_id: number;
  content: string;
  parent_id?: number;
  likes_count: number;
  created_at: string;
  updated_at: string;
  user?: User;
  replies?: Comment[];
}

export interface AuthResponse {
  user: User;
  access_token: string;
  refresh_token: string;
}

export interface VideoListResponse {
  videos: Video[];
  total_count: number;
  page: number;
  page_size: number;
}
