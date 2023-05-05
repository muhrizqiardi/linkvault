'use client';

import useSWR from 'swr';
import fetcher from './fetcher';

export function useUser() {
  const { data, isLoading, error } = useSWR(
    `/api/auth`,
    fetcher<{
      id: string;
      email: string;
      full_name: string;
      created_at: string;
      updated_at: string;
    }>,
  );

  return { data, isLoading, error };
}
