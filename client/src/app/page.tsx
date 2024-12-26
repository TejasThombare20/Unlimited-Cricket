import { ArrowRight } from "lucide-react";
import Link from "next/link";

export default function Home() {
  return (
    <main className="h-full ">
      {/* <Navigation /> */}
      <section className="h-full w-full pt-[200px] relative flex items-center justify-center flex-col gap-10">
        <div
          aria-label="true"
          className="absolute bottom-0 left-0  right-0 top-0 bg-[linear-gradient(to_right,#4f4f4f2e_1px,transparent_1px),linear-gradient(to_bottom,#4f4f4f2e_1px,transparent_1px)] bg-[size:4rem_64px] [mask-image:radial-gradient(ellipse_60%_50%_at_50%_0%,#000_70%,transparent_110%)] -z-10"
        />
        <div className=" bg-gradient-to-r from-primary to-secondary-foreground text-transparent bg-clip-text relative">
          <h1 className=" max-w-7xl text-8xl font-bold text-center sm:text-[100px] ">
          Endless Cricket Action, One Click Away!
          </h1>
        </div>
        <Link href={"/dashboard"}>
          <button className="inline-flex h-12 animate-background-shine items-center justify-center rounded-md border border-gray-800 bg-[linear-gradient(110deg,#000103,45%,#1e2631,55%,#000103)] bg-[length:200%_100%] px-6 font-medium text-gray-200 transition-colors focus:outline-none focus:ring-2 focus:ring-gray-400 focus:ring-offset-2 focus:ring-offset-gray-50 gap-3">
            <span>Land on Action</span>
            <ArrowRight size={20} />
          </button>
        </Link>
      </section>
    </main>
  );
}
