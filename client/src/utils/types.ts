export interface Video {
    title: string;
    description: string;
    published_at: string;
    thumbnail_url: string;
    created_at: string;
  }
  
export interface ApiResponse {
    data: Video[];
    total: number;
    page: number;
  }
  