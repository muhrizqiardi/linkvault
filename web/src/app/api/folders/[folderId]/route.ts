import * as z from 'zod';
import { folderEntitySchema } from '@/entities';
import { createFolderPopupFormDto, updateFolderDtoSchema } from '@/schemas';
import { env } from '@/utils/env';
import { cookies } from 'next/headers';

export async function PATCH(
  req: Request,
  params: { params: { folderId: string } },
) {
  try {
    const token = cookies().get('token');
    if (token === undefined) throw new Error();

    const requestBody = updateFolderDtoSchema.parse(await req.json());

    const response = await fetch(
      `${env.API_URL}/folders/${params.params.folderId}`,
      {
        method: 'PATCH',
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

    const updatedFolder = z
      .object({
        success: z.boolean(),
        message: z.string(),
        data: folderEntitySchema,
      })
      .parse(responseBody);

    return new Response(JSON.stringify(updatedFolder), {
      status: 200,
    });
  } catch (error) {
    console.error(error);
    return new Response(
      JSON.stringify({
        success: false,
        message: 'Failed to update user',
      }),
      {
        status: 500,
      },
    );
  }
}

export async function DELETE(
  req: Request,
  params: { params: { folderId: string } },
) {
  try {
    const token = cookies().get('token');
    if (token === undefined) throw new Error();

    const response = await fetch(
      `${env.API_URL}/folders/${params.params.folderId}`,
      {
        method: 'DELETE',
        headers: {
          accept: 'application/json',
          Authorization: `Bearer ${token.value}`,
        },
      },
    );
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
        message: 'Failed to delete user',
      }),
      {
        status: 500,
      },
    );
  }
}
