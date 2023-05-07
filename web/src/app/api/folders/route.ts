import * as z from 'zod';
import { env } from '@/utils/env';
import { cookies } from 'next/headers';
import { createFolderPopupFormDto } from '@/schemas';
import { folderEntitySchema } from '@/entities';

export async function POST(req: Request) {
  try {
    const token = cookies().get('token');
    if (token === undefined) throw new Error();

    const requestBody = createFolderPopupFormDto.parse(await req.json());

    const response = await fetch(`${env.API_URL}/folders`, {
      method: 'POST',
      body: JSON.stringify(requestBody),
      headers: {
        accept: 'application/json',
        Authorization: `Bearer ${token.value}`,
        'Content-Type': 'application/json',
      },
    });
    if (!response.ok) throw response;

    const responseBody = await response.json();

    const newFolder = z
      .object({
        success: z.boolean(),
        message: z.string(),
        data: folderEntitySchema,
      })
      .parse(responseBody);

    return new Response(JSON.stringify(newFolder), {
      status: 201,
    });
  } catch (error) {
    console.error(error);
    return new Response(
      JSON.stringify({
        success: false,
        message: 'Failed to create user',
      }),
      {
        status: 500,
      },
    );
  }
}
