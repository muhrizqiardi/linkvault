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

        if (!response.ok) throw new Error();

        const responseBody = await response.json();
        res
          .status(201)
          .setHeader(
            'Set-Cookie',
            `token=${responseBody.data}; expires=${
              new Date().getTime() + 1000 * 60 * 60 * 24 * 30 * 6
            }; Path=/`,
          )
          .send({
            success: true,
            message: 'Successfully signed in',
          });
      } catch (error) {
        res.status(401).send({ success: false, message: 'Failed to sign in' });
      }
      break;
    case 'GET':
      try {
        if (req.cookies['token'] === undefined || req.cookies['token'] === null)
          throw new Error('Unauthorized');

        const response = await fetch(`${env.API_URL}/auth`, {
          method: 'GET',
          headers: {
            Authorization: `Bearer ${req.cookies.token}`,
            'Content-Type': 'application/json',
          },
        });

        if (!response.ok) throw new Error();
        const responseBody = await response.json();

        res
          .status(200)
          .setHeader(
            'Set-Cookie',
            `token=${req.cookies['token']}; expires=${
              new Date().getTime() + 1000 * 60 * 60 * 24 * 30 * 6
            }; Path=/`,
          )
          .send(responseBody.data);
      } catch (error) {
        res
          .status(401)
          .setHeader(
            'Set-Cookie',
            'token=deleted; Path=/; expires=Thu, 01 Jan 1970 00:00:00 GMT',
          );
      }
      break;
    case 'DELETE':
      res
        .status(401)
        .setHeader(
          'Set-Cookie',
          'token=deleted; Path=/; expires=Thu, 01 Jan 1970 00:00:00 GMT',
        );
      break;

    default:
      res.status(200).send({
        success: true,
        message: 'Not Found',
      });
      break;
  }
}
