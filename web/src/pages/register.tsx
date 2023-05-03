import { RegisterPageForm } from '@/components/register-page-form';
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import Link from 'next/link';

export default function SignInPage() {
  return (
    <div className="h-screen pt-24">
      <Card className="w-[350px] mx-auto">
        <CardHeader>
          <CardTitle>Register account</CardTitle>
          <CardDescription>
            Already have an account? Create it{' '}
            <Link
              className="font-medium text-primary underline underline-offset-4"
              href="/sign-in"
            >
              here
            </Link>
            .
          </CardDescription>
        </CardHeader>
        <CardContent>
          <RegisterPageForm />
        </CardContent>
      </Card>
    </div>
  );
}
