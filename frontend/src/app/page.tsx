"use client"
import Link from "next/link"

const JobZipLanding = () => {
  return (
    <div className="relative isolate overflow-hidden">
      {/* Background Pattern */}
      <svg
        className="absolute inset-0 -z-10 h-full w-full stroke-gray-300 [mask-image:radial-gradient(100%_100%_at_top_right,white,transparent)]"
        aria-hidden="true"
      >
        <defs>
          <pattern
            id="983e3e4c-de6d-4c3f-8d64-b9761d1534cc"
            width="200"
            height="200"
            x="50%"
            y="-1"
            patternUnits="userSpaceOnUse"
          >
            <path d="M.5 200V.5H200" fill="none" />
          </pattern>
        </defs>
        <svg x="50%" y="-1" className="overflow-visible fill-gray-200">
          <path
            d="M-200 0h201v201h-201Z M600 0h201v201h-201Z M-400 600h201v201h-201Z M200 800h201v201h-201Z"
            strokeWidth="0"
          />
        </svg>
        <rect width="100%" height="100%" strokeWidth="0" fill="url(#983e3e4c-de6d-4c3f-8d64-b9761d1534cc)" />
      </svg>

      {/* Gradient Blob */}
      <div
        className="absolute left-[calc(50%-4rem)] top-10 -z-10 transform-gpu blur-3xl sm:left-[calc(50%-18rem)] lg:left-48 lg:top-[calc(50%-30rem)] xl:left-[calc(50%-24rem)]"
        aria-hidden="true"
      >
        <div
          className="aspect-[1108/632] w-[69.25rem] bg-gradient opacity-20"
          style={{
            clipPath:
              "polygon(73.6% 51.7%, 91.7% 11.8%, 100% 46.4%, 97.4% 82.2%, 92.5% 84.9%, 75.7% 64%, 55.3% 47.5%, 46.5% 49.4%, 45% 62.9%, 50.3% 87.2%, 21.3% 64.1%, 0.1% 100%, 5.4% 51.1%, 21.4% 63.9%, 58.9% 0.2%, 73.6% 51.7%)",
          }}
        />
      </div>

      {/* Header */}
      <header>
        <nav className="mx-auto flex max-w-7xl items-center justify-between p-6 lg:px-8" aria-label="Global">
          {/* Logo */}
          <div className="flex lg:flex-1">
            <Link href="/" className="-m-1.5 p-1.5 flex items-center">
              <span className="sr-only">Job.zip</span>
              <img
                className="h-10 w-auto"
                src="https://mhiqwtehsmorqwewxvqx.supabase.co/storage/v1/object/public/Jobzip/jobzip-lq-ts-w.png"
                alt="Job.zip logo"
              />
            </Link>
          </div>

          {/* Mobile Menu Button */}
          <div className="flex lg:hidden">
            <button
              type="button"
              className="mr-2 mt-1 btn btn-ghost btn-small"
            >
              <span className="sr-only">Search</span>
              <svg className="h-5 w-5" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 24 24">
                <path
                  fillRule="evenodd"
                  clipRule="evenodd"
                  d="M10.5 3.75a6.75 6.75 0 1 0 0 13.5 6.75 6.75 0 0 0 0-13.5ZM2.25 10.5a8.25 8.25 0 1 1 14.59 5.28l4.69 4.69a.75.75 0 1 1-1.06 1.06l-4.69-4.69A8.25 8.25 0 0 1 2.25 10.5Z"
                />
              </svg>
            </button>
            <button
              type="button"
              className="btn btn-ghost"
            >
              <span className="sr-only">Open main menu</span>
              <svg
                className="h-6 w-6"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
                strokeWidth="1.5"
              >
                <path strokeLinecap="round" strokeLinejoin="round" d="M3.75 6.75h16.5M3.75 12h16.5m-16.5 5.25h16.5" />
              </svg>
            </button>
          </div>

          {/* Desktop Navigation */}
          <div className="hidden lg:flex lg:gap-x-12">
            <a href="/category/all" className="nav-link">
              Trends
            </a>
            <a href="/hiring-category/all" className="nav-link">
              Companies
            </a>
            <a href="/about" className="nav-link">
              About
            </a>
            <button
              type="button"
              className="btn btn-ghost btn-small"
            >
              <span className="sr-only">Search</span>
              <svg className="h-5 w-5" xmlns="http://www.w3.org/2000/svg" fill="currentColor" viewBox="0 0 24 24">
                <path
                  fillRule="evenodd"
                  clipRule="evenodd"
                  d="M10.5 3.75a6.75 6.75 0 1 0 0 13.5 6.75 6.75 0 0 0 0-13.5ZM2.25 10.5a8.25 8.25 0 1 1 14.59 5.28l4.69 4.69a.75.75 0 1 1-1.06 1.06l-4.69-4.69A8.25 8.25 0 0 1 2.25 10.5Z"
                />
              </svg>
            </button>
          </div>

          {/* CTA Buttons */}
          <div className="hidden lg:flex lg:flex-1 lg:justify-end gap-4">
            <a
              href="https://mail.job.zip"
              target="_blank"
              rel="noopener noreferrer"
              className="hidden xl:block btn btn-secondary"
            >
              Trends Newsletter
            </a>
            <a
              href="https://mail.job.zip"
              target="_blank"
              rel="noopener noreferrer"
              className="block xl:hidden btn btn-secondary"
            >
              Newsletter
            </a>
            <a
              href="https://fantastic.jobs/"
              target="_blank"
              rel="noopener noreferrer"
              className="btn btn-primary"
            >
              Job Data by Fantastic.jobs
            </a>
          </div>
        </nav>
      </header>

      {/* Main Content */}
      <main className="container mx-auto flex min-h-screen flex-col items-center">
        <div className="max-w-4xl px-6 pb-12 pt-8 sm:pb-40 text-center relative">
          <h1 className="mt-24 sm:mt-16 lg:mt-8 heading-1 animate-fade-in-up">
            We analyze 10 million jobs per month to discover trends in tech
          </h1>
          
          <p className="paragraph paragraph-center animate-fade-in-up" style={{ animationDelay: '0.2s' }}>
            Every month we send out a newsletter with the latest trends.
          </p>

          <div className="flex justify-center mt-8 relative z-10 animate-fade-in-up" style={{ animationDelay: '0.4s' }}>
            <div className="search-container">
              <div className="relative">
                <div className="search-icon">
                  <svg
                    className="h-5 w-5"
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth="2"
                      d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
                    />
                  </svg>
                </div>
                <input
                  type="text"
                  className="input input-with-icon"
                  placeholder="Search by company, role, or skill..."
                />
              </div>
            </div>
          </div>
        </div>
      </main>
    </div>
  )
}

export default JobZipLanding