import { SignInPageFormDto, signInPageFormDtoSchema } from '@/schemas';
import useError from '@/utils/use-error';
import { zodResolver } from '@hookform/resolvers/zod';
import { Loader2 } from 'lucide-react';
import { useRouter } from 'next/router';
import { useEffect, useState } from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';
import { Button } from './ui/button';
import { Input } from './ui/input';
import { Label } from './ui/label';

export function SignInPageForm() {
  const router = useRouter();
  const [isLoading, setIsLoading] = useState(false);
  const { isError, setError } = useError(false, '');
  const { register, handleSubmit } = useForm<SignInPageFormDto>({
    resolver: zodResolver(signInPageFormDtoSchema),
  });
  const [isClientSide, setIsClientSide] = useState(false);

  const onSubmit: SubmitHandler<SignInPageFormDto> = async (data) => {
    setIsLoading(true);
    try {
      const response = await fetch('/api/auth', {
        method: 'POST',
        body: JSON.stringify(data),
      });

      if (!response.ok) throw new Error();
      router.push('/');
    } catch (error) {
      setError(true, 'Failed to sign in');
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    setIsClientSide(true);
  });

  if (isClientSide)
    return (
      <form onSubmit={handleSubmit(onSubmit)}>
        <div className="grid w-full items-center gap-4">
          <div className="flex flex-col space-y-1.5">
            <Label htmlFor="name">Email</Label>
            <Input
              id="name"
              type="email"
              placeholder="user@example.com"
              {...register('email')}
            />
          </div>
          <div className="flex flex-col space-y-1.5">
            <Label htmlFor="name">Password</Label>
            <Input
              id="name"
              type="password"
              placeholder="Enter your password here"
              {...register('password')}
            />
          </div>
          <Button disabled={isLoading} type="submit" className="mt-4">
            {isLoading ? (
              <Loader2 className="mr-2 h-4 w-4 animate-spin" />
            ) : null}
            Sign in
          </Button>
        </div>
      </form>
    );

  return null;
}
