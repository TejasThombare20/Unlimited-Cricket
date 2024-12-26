"use client";
import { Loader2 } from "lucide-react";
import Image from "next/image";
import React, { useEffect, useState } from "react";
import { useInView } from "react-intersection-observer";
import ApiHandler from "@/services/api-services";
import axios from "axios";
import ErrorMessage from "@/components/ErrorState";
import EmptyState from "@/components/EmptyState";
import { formatDate, getTimeAgo } from "@/utils/getTimefunc";
import { ApiResponse, Video } from "@/utils/types";

const page = () => {
  const [videos, setVideos] = useState<Video[]>([]);
  const [page, setPage] = useState(1);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [hasMore, setHasMore] = useState(true);
  const { ref, inView } = useInView();

  const fetchVideos = async (pageNum: number) => {
    try {
      setLoading(true);
      setError(null);

      const response = await ApiHandler.get<ApiResponse>(
        `/api/v1/lists?page=${pageNum}&page_size=20`
      );

      if (pageNum === 1) {
        setVideos(response.data);
      } else {
        setVideos((prev) => [...prev, ...response.data]);
      }

      setHasMore(response.data.length === 20);
    } catch (error) {
      if (axios.isAxiosError(error)) {
        setError(error.response?.data?.message || "Failed to fetch videos");
      } else {
        setError("An unexpected error occurred");
      }
    } finally {
      setLoading(false);
    }
  };

  // Effect hooks remain the same
  useEffect(() => {
    fetchVideos(1);
  }, []);

  useEffect(() => {
    if (inView && !loading && hasMore && !error) {
      setPage((prev) => {
        const nextPage = prev + 1;
        fetchVideos(nextPage);
        return nextPage;
      });
    }
  }, [inView, loading, hasMore, error]);

  return (
    <div className="container mx-auto px-4 py-8">
      <main className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {error ? (
          <ErrorMessage message={error} />
        ) : videos?.length === 0 && !loading ? (
          <EmptyState />
        ) : (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {videos.map((video, index) => (
              <div
                key={`${video.title}-${index}`}
                className="bg-white dark:bg-gray-800 rounded-lg overflow-hidden shadow-md hover:shadow-lg transition-shadow duration-300"
              >
                <div className="relative aspect-video">
                  <Image
                    src={video?.thumbnail_url}
                    alt={video?.title}
                    quality={100}
                    className="object-cover"
                    priority={index < 6}
                    blurDataURL="data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAAAQABAAD/4gHYSUNDX1BST0ZJTEUAAQEAAAHIAAAAAAQwAABtbnRyUkdCIFhZWiAH4AABAAEAAAAAAABhY3NwAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAA9tYAAQAAAADTLQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAlkZXNjAAAA8AAAACRyWFlaAAABFAAAABRnWFlaAAABKAAAABRiWFlaAAABPAAAABR3dHB0AAABUAAAABRyVFJDAAABZAAAAChnVFJDAAABZAAAAChiVFJDAAABZAAAAChjcHJ0AAABjAAAADxtbHVjAAAAAAAAAAEAAAAMZW5VUwAAAAgAAAAcAHMAUgBHAEJYWVogAAAAAAAAb6IAADj1AAADkFhZWiAAAAAAAABimQAAt4UAABjaWFlaIAAAAAAAACSgAAAPhAAAts9YWVogAAAAAAAA9tYAAQAAAADTLXBhcmEAAAAAAAQAAAACZmYAAPKnAAANWQAAE9AAAApbAAAAAAAAAABtbHVjAAAAAAAAAAEAAAAMZW5VUwAAACAAAAAcAEcAbwBvAGcAbABlACAASQBuAGMALgAgADIAMAAxADb/2wBDABQODxIPDRQSEBIXFRQdHx0cHBwcHx0cHBwcHBwcHBwcHBwcHBwcHBwcHBwcHBwcHBwcHBwcHBwcHBwcHBwcHBz/2wBDAR0XFx8bHxwcHBwfHBwcHBwcHBwcHBwcHBwcHBwcHBwcHBwcHBwcHBwcHBwcHBwcHBwcHBwcHBwcHBwcHBwcHBz/wAARCAAIAAoDASIAAhEBAxEB/8QAFQABAQAAAAAAAAAAAAAAAAAAAAb/xAAUEAEAAAAAAAAAAAAAAAAAAAAA/8QAFQEBAQAAAAAAAAAAAAAAAAAAAAX/xAAUEQEAAAAAAAAAAAAAAAAAAAAA/9oADAMBAAIRAxEAPwCdABmX/9k="
                    width={640}
                    height={360}
                  />
                </div>
                <div className="p-4">
                  <h3 className="font-semibold text-lg line-clamp-2 mb-2 text-gray-900 dark:text-gray-100">
                    {video?.title}
                  </h3>
                  <div className="flex justify-between items-center text-sm text-gray-500 dark:text-gray-400 mb-3">
                    <span>{formatDate(video?.published_at)}</span>
                    <span>{getTimeAgo(video?.published_at)}</span>
                  </div>
                  <p className="text-sm text-gray-600 dark:text-gray-300 line-clamp-3">
                    {video?.description}
                  </p>
                </div>
              </div>
            ))}
          </div>
        )}

        <div ref={ref} className="w-full py-8 flex justify-center">
          {loading && (
            <div className="flex items-center gap-2 text-gray-600 dark:text-gray-300">
              <Loader2 className="h-6 w-6 animate-spin" />
              <span>Loading more videos...</span>
            </div>
          )}
        </div>
      </main>
    </div>
  );
};

export default page;
