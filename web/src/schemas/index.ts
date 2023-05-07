import { z } from 'zod';

export const registerPageFormDtoSchema = z.object({
  full_name: z.string(),
  email: z.string().email(),
  password: z.string().length(8),
  confirm_password: z.string().length(8),
});

export interface RegisterPageFormDto
  extends z.infer<typeof registerPageFormDtoSchema> {}

export const signInPageFormDtoSchema = z.object({
  email: z.string().email(),
  password: z.string().length(8),
});

export interface SignInPageFormDto
  extends z.infer<typeof signInPageFormDtoSchema> {}

export const createFolderPopupFormDto = z.object({
  name: z.string(),
});

export interface CreateFolderPopupFormDto
  extends z.infer<typeof createFolderPopupFormDto> {}

export const updateFolderDtoSchema = createFolderPopupFormDto.partial();

export interface UpdateFolderDto
  extends z.infer<typeof updateFolderDtoSchema> {}
