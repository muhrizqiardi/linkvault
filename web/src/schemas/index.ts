import { z } from 'zod';

export const registerPageFormDtoSchema = z.object({
  full_name: z.string(),
  email: z.string().email(),
  password: z.string().length(8),
  confirm_password: z.string().length(8),
});
// .refine((arg) => arg.password !== arg.confirm_password, {
//   message: 'Passwords do not match',
// });

export interface RegisterPageFormDto
  extends z.infer<typeof registerPageFormDtoSchema> {}
