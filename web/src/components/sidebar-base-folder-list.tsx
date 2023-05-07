import { folderEntitySchema } from '@/entities';
import { HiEllipsisVertical, HiFolder, HiFolderPlus } from 'react-icons/hi2';
import { Button, buttonVariants } from './ui/button';
import { Skeleton } from './ui/skeleton';
import { cookies } from 'next/headers';
import { env } from '@/utils/env';
import * as z from 'zod';
import Link from 'next/link';

const getFolders = async () => {
  try {
    const token = cookies().get('token');
    if (token === undefined) throw new Error();

    const response = await fetch(
      `${env.API_URL}/folders?orderBy=updatedAt_DESC&limit=10&page=1`,
      {
        method: 'GET',
        headers: {
          Accept: 'application/json',
          Authorization: `Bearer ${token.value}`,
        },
      },
    );
    if (!response.ok) throw new Error('Request failed');

    const responseBody = await response.json();
    const responseBodySchema = z.object({
      success: z.boolean(),
      message: z.string(),
      data: z.array(folderEntitySchema),
    });
    const { data } = responseBodySchema.parse(responseBody);

    return data;
  } catch (error) {
    console.error(error);
    return null;
  }
};

export default async function SidebarBaseFolderList() {
  const folders = await getFolders();

  if (folders !== null)
    return (
      <div className="px-4 py-2">
        <div className="flex justify-between items-center">
          <h2 className="px-2 text-lg font-semibold tracking-tight">Folders</h2>
          <div className="flex gap-1">
            <Button variant="ghost" size="sm" className="rounded-full h-9 w-9">
              <HiEllipsisVertical />
            </Button>
            <Button variant="ghost" size="sm" className="rounded-full h-9 w-9">
              <HiFolderPlus />
            </Button>
          </div>
        </div>
        {folders.length === 0 ? (
          <div className="py-2 px-5 border-2 border-dashed rounded-xl h-24 flex flex-col items-center text-center justify-center">
            <p className="text-sm">You haven't created any folder.</p>
          </div>
        ) : null}
        <div className="space-y-1 py-2 flex flex-col">
          {folders.map((folder, index) => (
            <Link
              key={index}
              href={`/folders/${folder.id}`}
              className={
                buttonVariants({
                  variant: 'ghost',
                  size: 'sm',
                  justify: 'start',
                }) + 'w-full'
              }
            >
              <HiFolder className="mr-2 h-4 w-4" />
              {folder.name}
            </Link>
          ))}
        </div>
      </div>
    );

  return null;
}

export const SidebarBaseFolderListSkeleton = () => (
  <div className="px-4 py-2">
    <div className="flex justify-between items-center">
      <h2 className="px-2 text-lg font-semibold tracking-tight">Folders</h2>
      <div className="flex gap-1">
        <Button variant="outline" size="sm">
          <HiEllipsisVertical />
        </Button>
        <Button variant="outline" size="sm">
          <HiFolderPlus />
        </Button>
      </div>
    </div>
    <div className="space-y-1 py-2">
      <div className="h-9 px-3 flex items-center">
        <Skeleton className="h-4 w-4 mr-2" />
        <Skeleton className="h-4 w-full" />
      </div>
      <div className="h-9 px-3 flex items-center">
        <Skeleton className="h-4 w-4 mr-2" />
        <Skeleton className="h-4 w-full" />
      </div>
      <div className="h-9 px-3 flex items-center">
        <Skeleton className="h-4 w-4 mr-2" />
        <Skeleton className="h-4 w-full" />
      </div>
      <div className="h-9 px-3 flex items-center">
        <Skeleton className="h-4 w-4 mr-2" />
        <Skeleton className="h-4 w-full" />
      </div>
    </div>
  </div>
);
