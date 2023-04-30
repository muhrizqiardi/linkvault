import { Button } from '@/components/ui/button';
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import { Input } from '@/components/ui/input';
import { Label } from '@/components/ui/label';
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
          <form>
            <div className="grid w-full items-center gap-4">
              <div className="flex flex-col space-y-1.5">
                <Label htmlFor="name">Full name</Label>
                <Input id="name" type="email" placeholder="Doe John" />
              </div>
              <div className="flex flex-col space-y-1.5">
                <Label htmlFor="name">Email</Label>
                <Input id="name" type="email" placeholder="user@example.com" />
              </div>
              <div className="flex flex-col space-y-1.5">
                <Label htmlFor="name">Password</Label>
                <Input
                  id="name"
                  type="password"
                  placeholder="Enter your password here"
                />
              </div>
              <div className="flex flex-col space-y-1.5">
                <Label htmlFor="name">Password</Label>
                <Input
                  id="name"
                  type="password"
                  placeholder="Confirm your password"
                />
              </div>
            </div>
          </form>
        </CardContent>
        <CardFooter className="flex justify-end">
          <Button>Create account</Button>
        </CardFooter>
      </Card>
    </div>
  );
}
