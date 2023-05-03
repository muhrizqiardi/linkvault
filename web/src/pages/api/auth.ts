import { signInPageFormDtoSchema } from '@/schemas';
import { env } from '@/utils/env';
import { NextApiRequest, NextApiResponse } from 'next';

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse,
) {
  switch (req.method) {
    case 'POST':
      try {
        const body = signInPageFormDtoSchema.parse(JSON.parse(req.body));
        const response = await fetch(`${env.API_URL}/auth`, {
          method: 'POST',
          body: JSON.stringify(body),
          headers: {
            'Content-Type': 'application/json',
          },
        });
        const responseBody = await response.json();
        return res
          .status(201)
          .setHeader(
            'Set-Cookie',
            `token=${responseBody.data}; expires=${
              new Date().getTime() + 1000 * 60 * 60 * 24 * 30 * 6
            }`,
          )
          .send({
            success: true,
            message: 'Successfully signed in',
          });
      } catch (error) {
        return res
          .status(401)
          .send({ success: false, message: 'Failed to sign in' });
      }
    case 'GET':
      try {
        if (!req.cookies['token']) throw new Error('Unauthorized');
        // TODO: fetch GET /auth to check the validity of the token received
        return res
          .status(200)
          .setHeader(
            'Set-Cookie',
            `token=${req.cookies['token']}; expires=${
              new Date().getTime() + 1000 * 60 * 60 * 24 * 30 * 6
            }`,
          );
      } catch (error) {
        return res
          .status(401)
          .setHeader(
            'Set-Cookie',
            'token=deleted; path=/; expires=Thu, 01 Jan 1970 00:00:00 GMT',
          );
      }
    case 'DELETE':
      return res
        .status(401)
        .setHeader(
          'Set-Cookie',
          'token=deleted; path=/; expires=Thu, 01 Jan 1970 00:00:00 GMT',
        );

    default:
      return res.status(200).send({
        success: true,
        message: 'Not Found',
      });
  }
}
