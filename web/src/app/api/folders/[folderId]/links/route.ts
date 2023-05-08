import { createNewLinkDtoSchema, linkEntitySchema } from '@/entities';
import { env } from '@/utils/env';
import { cookies } from 'next/headers';
import * as z from 'zod';

export async function POST(
  req: Request,
  params: { params: { folderId: string } },
) {
  try {
    const token = cookies().get('token');
    if (token === undefined) throw new Error();

    const requestBody = createNewLinkDtoSchema.parse(await req.json());

    const response = await fetch(
      `${env.API_URL}/folders/${params.params.folderId}/links`,
      {
        method: 'POST',
        body: JSON.stringify(requestBody),
        headers: {
          accept: 'application/json',
          Authorization: `Bearer ${token.value}`,
          'Content-Type': 'application/json',
        },
      },
    );
    if (!response.ok) throw response;

    const responseBody = await response.json();

    const newLink = z
      .object({
        success: z.boolean(),
        message: z.string(),
        data: linkEntitySchema,
      })
      .parse(responseBody);

    return new Response(JSON.stringify(newLink), {
      status: 201,
    });
  } catch (error) {
    console.error(error);
    return new Response(
      JSON.stringify({
        success: false,
        message: 'Failed to create link',
      }),
      {
        status: 500,
      },
    );
  }
}
