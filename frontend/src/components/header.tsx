import Link from "next/link"
import { Button } from "@/components/ui/button"

export default function Header() {
  return (
    <header>
      <nav className="mx-auto flex max-w-7xl items-center justify-between p-6 lg:px-8" aria-label="Global">
        <div className="flex lg:flex-1">
          <Link href="/" className="-m-1.5 p-1.5 flex items-center">
            <span className="sr-only">JobZip</span>
            <img
              className="h-10 w-auto"
              src="https://mhiqwtehsmorqwewxvqx.supabase.co/storage/v1/object/public/Jobzip/jobzip-lq-ts-w.png"
              alt="JobZip logo"
            />
          </Link>
        </div>

        <div className="flex lg:hidden">
          <button type="button" className="btn btn-ghost">
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

        <div className="hidden lg:flex lg:gap-x-12">
          <a href="/jobs" className="nav-link">
            Jobs
          </a>
          <a href="/companies" className="nav-link">
            Companies
          </a>
          <a href="/services" className="nav-link">
            Services
          </a>
        </div>

        <div className="hidden lg:flex lg:flex-1 lg:justify-end gap-4">
          <Button variant="outline" asChild>
            <Link href="/login">Login</Link>
          </Button>
          <Button asChild className="bg-orange-500 hover:bg-orange-600">
            <Link href="/register">Register</Link>
          </Button>
          <Button variant="ghost" className="text-gray-500" asChild>
            <Link href="/employers">For employers</Link>
          </Button>
        </div>
      </nav>
    </header>
  )
}
