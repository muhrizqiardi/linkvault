import { GetServerSideProps } from 'next';

export const getServerSideProps: GetServerSideProps = async (context) => {
  context.res.setHeader(
    'Set-Cookie',
    'token=deleted; Path=/; expires=Thu, 01 Jan 1970 00:00:00 GMT',
  );

  return {
    redirect: {
      destination: '/sign-in',
    },
    props: {},
  };
};

export default function SignOutPage() {
  return null;
}
