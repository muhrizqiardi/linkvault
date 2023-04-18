import Image from 'next/image';
import {
  HiArrowTopRightOnSquare,
  HiPencil,
  HiShare,
  HiTrash,
} from 'react-icons/hi2';

export default function ListViewLinkItem() {
  return (
    <article className="card card-compact bg-base-100 rounded-none border-b border-b-base-200 last:border-b-0">
      <div className="card-body">
        <div className="flex gap-2 md:gap-4">
          <div className="flex-shrink-0">
            <div className="relative aspect-square w-20">
              <Image
                src="https://picsum.photos/id/237/200/300"
                fill
                alt=""
                className="object-cover"
              />
            </div>
          </div>
          <div className="flex flex-col">
            <p>twitter.com · 2h ago</p>
            <h2 className="card-title">
              Lorem, ipsum dolor sit amet consectetur adipisicing elit. Dolorum
              iusto facilis quo. Atque nam illum vel. Velit, repellendus sequi
              ipsam debitis totam obcaecati autem sapiente iste asperiores
              soluta laboriosam molestias.
            </h2>
            <p>If a dog chews shoes whose shoes does he choose?</p>
          </div>
        </div>
        <div className="card-actions justify-end">
          <button className="btn btn-sm btn-ghost btn-square">
            <HiTrash />
          </button>
          <button className="btn btn-sm btn-ghost btn-square">
            <HiPencil />
          </button>
          <button className="btn btn-sm btn-ghost btn-square">
            <HiArrowTopRightOnSquare />
          </button>
          <button className="btn btn-sm btn-ghost btn-square">
            <HiShare />
          </button>
        </div>
      </div>
    </article>
  );
}
