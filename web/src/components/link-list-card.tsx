import { HiFolder } from 'react-icons/hi2';
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';

export function LinkListCard() {
  return (
    <Card className="border-t-0 border-r-0 border-l-0 shadow-none rounded-none">
      <div className="flex">
        <div className="p-3.5 pr-0">
          <img
            src="https://picsum.photos/id/237/200/300"
            alt=""
            className="aspect-[1200/628] w-full object-contain bg-neutral-200 rounded-sm"
          />
        </div>

        <div className="block">
          <CardHeader>
            <CardDescription className="inline-flex items-center">
              twitter.com · 2h ago ·{' '}
              <span className="ml-2 inline-flex items-center">
                <HiFolder className="mr-1" />
                Bookmarked Tweets
              </span>
            </CardDescription>
            <CardTitle>Kabar Penumpang on Twitter</CardTitle>
          </CardHeader>
          <CardContent className="text-sm">
            <p>
              Menolak Lupa Proposal Kereta Cepat Jakarta Bandung Jepang: US$6,2
              miliar (75 persennya ditanggung Jepang berupa pinjaman tenor 40
              thn dan bunga 0,1 persen). Menurut kajian, proyek kereta cepat
              sulit dgn skema b to b atau tanpa jaminan pemerintah. China pun
              dtg skema menggiurkan
            </p>
          </CardContent>
        </div>
      </div>
    </Card>
  );
}
