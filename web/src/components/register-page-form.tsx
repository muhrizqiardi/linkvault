import { RegisterPageFormDto, registerPageFormDtoSchema } from '@/schemas';
import { useState, useEffect } from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';
import { Button } from './ui/button';
import { Input } from './ui/input';
import { Label } from './ui/label';
import { zodResolver } from '@hookform/resolvers/zod';
import useError from '@/utils/use-error';
import { Loader2 } from 'lucide-react';
import { useRouter } from 'next/router';

export function RegisterPageForm() {
  const [isLoading, setIsLoading] = useState(false);
  const { isError, setError } = useError(false);
  const router = useRouter();
  const [isClientSide, setIsClientSide] = useState(false);
  const { register, handleSubmit } = useForm<RegisterPageFormDto>({
    resolver: zodResolver(registerPageFormDtoSchema),
  });

  const onSubmit: SubmitHandler<RegisterPageFormDto> = async (data) => {
    setIsLoading(true);
    try {
      await fetch('/api/users', {
        method: 'POST',
        body: JSON.stringify(data),
      });

      await fetch('/api/auth', {
        method: 'POST',
        body: JSON.stringify({
          email: data.email,
          password: data.password,
        }),
      });

      router.push('/');
    } catch (error) {
      setError(true, 'Failed to register user');
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
            <Label htmlFor="full_name">Full name</Label>
            <Input
              id="full_name"
              type="text"
              placeholder="Doe John"
              {...register('full_name')}
            />
          </div>
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
          <div className="flex flex-col space-y-1.5">
            <Label htmlFor="name">Confirm Password</Label>
            <Input
              id="name"
              type="password"
              placeholder="Confirm your password"
              {...register('confirm_password')}
            />
          </div>
          <Button disabled={isLoading} type="submit" className="mt-4">
            {isLoading ? (
              <Loader2 className="mr-2 h-4 w-4 animate-spin" />
            ) : null}
            Create account
          </Button>
        </div>
      </form>
    );

  return null;
}
