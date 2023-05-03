import { SignInPageForm } from '@/components/sign-in-page-form';
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import Link from 'next/link';

export default function SignInPage() {
  return (
    <div className="h-screen pt-24">
      <Card className="w-[350px] mx-auto">
        <CardHeader>
          <CardTitle>Sign in to account</CardTitle>
          <CardDescription>
            Don't have an account? Sign in{' '}
            <Link
              className="font-medium text-primary underline underline-offset-4"
              href="/register"
            >
              here
            </Link>
            .
          </CardDescription>
        </CardHeader>
        <CardContent>
          <SignInPageForm />
        </CardContent>
        <CardFooter className="flex justify-end"></CardFooter>
      </Card>
    </div>
  );
}
