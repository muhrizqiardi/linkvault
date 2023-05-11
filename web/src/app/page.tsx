import * as z from 'zod';
import { LinkListCard } from '@/components/link-list-card';
import { TopNavBase } from '@/components/top-nav-base';
import { env } from '@/utils/env';
import { cookies } from 'next/headers';
import { linkEntitySchema } from '@/entities';

const getLinks = async () => {
  try {
    const token = cookies().get('token');
    if (token === undefined) throw new Error();

    const response = await fetch(
      `${env.API_URL}/links?orderBy=updatedAt_DESC&limit=10&page=1`,
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

export default async function HomePage() {
  const links = await getLinks();

  return (
    <>
      <div className="h-14">
        <TopNavBase />
      </div>
      <div className="flex flex-col">
        {links !== null ? (
          <div className="flex flex-col">
            {links.map((link, index) => (
              <LinkListCard link={link} key={index} />
            ))}
          </div>
        ) : null}
      </div>
    </>
  );
}
