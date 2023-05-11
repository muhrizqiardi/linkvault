import * as z from 'zod';
import { LinkListCard } from '@/components/link-list-card';
import { TopNavFolderPage } from '@/components/top-nav-folder-page';
import { TopNavFolderPageSkeleton } from '@/components/top-nav-folder-page-skeleton';
import { env } from '@/utils/env';
import { cookies } from 'next/headers';
import { folderEntitySchema, linkEntitySchema } from '@/entities';

const getFolderDetail = async (folderId: string) => {
  try {
    const token = cookies().get('token');
    if (token === undefined) throw new Error();

    const response = await fetch(`${env.API_URL}/folders/${folderId}`, {
      method: 'GET',
      headers: {
        Accept: 'application/json',
        Authorization: `Bearer ${token.value}`,
      },
    });
    if (!response.ok) throw new Error();
    const responseBody = await response.json();
    const folderDetail = z
      .object({
        success: z.boolean(),
        message: z.string(),
        data: folderEntitySchema,
      })
      .parse(responseBody);

    return folderDetail.data;
  } catch (error) {
    console.error(error);
    return null;
  }
};

const getLinksInFolder = async (folderId: string) => {
  try {
    const token = cookies().get('token');
    if (token === undefined) throw new Error();

    const response = await fetch(
      `${env.API_URL}/folders/${folderId}/links?orderBy=updatedAt_DESC&limit=10&page=1`,
      {
        method: 'GET',
        headers: {
          Accept: 'application/json',
          Authorization: `Bearer ${token.value}`,
        },
      },
    );
    if (!response.ok) throw new Error();
    const responseBody = await response.json();
    const links = z
      .object({
        success: z.boolean(),
        message: z.string(),
        data: z.array(linkEntitySchema),
      })
      .parse(responseBody);

    return links.data;
  } catch (error) {
    console.error(error);
    return null;
  }
};

export default async function FolderPage(props: {
  params: {
    folderId: string;
  };
}) {
  const links = await getLinksInFolder(props.params.folderId);
  const folderDetail = await getFolderDetail(props.params.folderId);

  return (
    <>
      <div className="h-14 border-b">
        {folderDetail !== null ? (
          <TopNavFolderPage folderDetail={folderDetail} />
        ) : (
          <TopNavFolderPageSkeleton />
        )}
      </div>
      {links !== null ? (
        <div className="flex flex-col">
          {links.map((link, index) => (
            <LinkListCard link={link} key={index} />
          ))}
        </div>
      ) : null}
    </>
  );
}
