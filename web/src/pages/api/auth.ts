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
          .setHeader('Set-Cookie', `token=${responseBody.data}`)
          .send({
            success: true,
            message: 'Successfully signed in',
          });
      } catch (error) {}
  }
}
