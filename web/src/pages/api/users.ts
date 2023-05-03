import { RegisterPageFormDto, registerPageFormDtoSchema } from '@/schemas';
import { env } from '@/utils/env';
import type { NextApiRequest, NextApiResponse } from 'next';
import z from 'zod';

type CreateUserResponse = {
  success: boolean;
  message: string;
  data?: {
    created_at: string;
    email: string;
    full_name: string;
    id: string;
    updated_at: string;
  };
};

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<CreateUserResponse>,
) {
  let body: RegisterPageFormDto;

  switch (req.method) {
    case 'POST':
      try {
        body = registerPageFormDtoSchema.parse(JSON.parse(req.body));
        const response = await fetch(`${env.API_URL}/users`, {
          method: 'POST',
          body: JSON.stringify({
            full_name: body.full_name,
            password: body.password,
            email: body.email,
          }),
          headers: {
            'Content-Type': 'application/json',
          },
        });
        const responseBody = await response.json();
        console.log(responseBody);
        return res.status(201).send({
          success: true,
          message: 'User registered',
          data: responseBody.data,
        });
      } catch (error) {
        if (error instanceof Error)
          return res.status(500).send({
            success: false,
            message: error.message,
          });
        return res.status(500).send({
          success: false,
          message: 'Internal Server Error',
        });
      }
    default:
      return res.status(404).send({
        success: false,
        message: 'Not Found',
      });
  }
}
