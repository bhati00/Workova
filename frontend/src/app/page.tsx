"use client"
import Link from "next/link"
import { useState } from "react"
import { Button } from "@/components/ui/button"
import { Combobox } from "@/components/ui/combobox"
import SearchBar, { type SearchItem } from "@/components/ui/searchBar"

// Dummy data for filters
const categories = [
  { value: "all", label: "All Categories" },
  { value: "technology", label: "Technology" },
  { value: "marketing", label: "Marketing" },
  { value: "design", label: "Design" },
  { value: "finance", label: "Finance" },
  { value: "sales", label: "Sales" },
  { value: "healthcare", label: "Healthcare" },
  { value: "education", label: "Education" },
]

const experienceOptions = [
  { value: "all", label: "Select experience" },
  { value: "fresher", label: "Fresher (less than 1 year)" },
  { value: "1year", label: "1 year" },
  { value: "2years", label: "2 years" },
  { value: "3years", label: "3 years" },
  { value: "4years", label: "4 years" },
  { value: "5years", label: "5 years" },
]

const skillsOptions = [
  { value: "all", label: "All Skills" },
  { value: "react", label: "React.js" },
  { value: "nodejs", label: "Node.js" },
  { value: "python", label: "Python" },
  { value: "javascript", label: "JavaScript" },
  { value: "typescript", label: "TypeScript" },
  { value: "aws", label: "AWS" },
  { value: "docker", label: "Docker" },
  { value: "kubernetes", label: "Kubernetes" },
]

const locations = [
  { id: 1, title: "New York, NY", subtitle: "United States" },
  { id: 2, title: "San Francisco, CA", subtitle: "United States" },
  { id: 3, title: "London", subtitle: "United Kingdom" },
  { id: 4, title: "Toronto", subtitle: "Canada" },
  { id: 5, title: "Berlin", subtitle: "Germany" },
  { id: 6, title: "Singapore", subtitle: "Singapore" },
  { id: 7, title: "Mumbai", subtitle: "India" },
  { id: 8, title: "Sydney", subtitle: "Australia" },
  { id: 9, title: "Remote", subtitle: "Work from anywhere" },
  { id: 10, title: "Hybrid", subtitle: "Remote + Office" },
]

const JobZipLanding = () => {
  const [selectedCategory, setSelectedCategory] = useState("")
  const [selectedExperience, setSelectedExperience] = useState("")
  const [selectedSkills, setSelectedSkills] = useState("")

  const handleLocationSelect = (location: SearchItem) => {
    console.log("Selected location:", location)
  }

  const handleSearch = () => {
    console.log("Search clicked with:", {
      category: selectedCategory,
      experience: selectedExperience,
      skills: selectedSkills,
    })
  }

  return (
    <div className="relative isolate overflow-hidden">
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

      <main className="container mx-auto flex min-h-screen flex-col items-center">
        <div className="max-w-6xl px-6 pb-12 pt-8 sm:pb-40 text-center relative w-full">
          <h1 className="mt-24 sm:mt-16 lg:mt-8 heading-1 animate-fade-in-up">Find your dream job now</h1>

          <p className="paragraph paragraph-center animate-fade-in-up" style={{ animationDelay: "0.2s" }}>
            5 lakh+ jobs for you to explore
          </p>

          <div className="mt-12 animate-fade-in-up" style={{ animationDelay: "0.4s" }}>
            <div className="bg-white rounded-2xl shadow-xl border border-gray-100 p-6 max-w-6xl mx-auto">
              <div className="hidden lg:block">
                <div className="grid grid-cols-5 gap-4 items-center">
                  <div>
                    <SearchBar
                      data={[
                        { id: "cloud", title: "Cloud, React.js," },
                        { id: "frontend", title: "Frontend Developer" },
                        { id: "backend", title: "Backend Developer" },
                        { id: "fullstack", title: "Full Stack Developer" },
                        { id: "devops", title: "DevOps Engineer" },
                      ]}
                      placeholder="Cloud, React.js,"
                      onSelect={(item: SearchItem) => console.log("Selected:", item)}
                      maxResults={5}
                      className="w-full"
                    />
                  </div>

                  <div>
                    <Combobox
                      items={experienceOptions}
                      value={selectedExperience}
                      onValueChange={setSelectedExperience}
                      placeholder="Select experience"
                      className="w-full h-12 bg-white border-gray-200 rounded-lg"
                    />
                  </div>

                  <div>
                    <Combobox
                      items={skillsOptions}
                      value={selectedSkills}
                      onValueChange={setSelectedSkills}
                      placeholder="Select skills"
                      className="w-full h-12 bg-white border-gray-200 rounded-lg"
                    />
                  </div>

                  <div>
                    <SearchBar
                      data={locations}
                      placeholder="Enter location"
                      onSelect={handleLocationSelect}
                      maxResults={5}
                      className="w-full"
                    />
                  </div>

                  <div>
                    <Button
                      onClick={handleSearch}
                      className="w-full h-12 bg-blue-600 hover:bg-blue-700 text-white font-semibold text-base rounded-lg shadow-md hover:shadow-lg transition-all duration-200"
                    >
                      Search
                    </Button>
                  </div>
                </div>
              </div>

              <div className="hidden md:block lg:hidden">
                <div className="space-y-4">
                  <div className="grid grid-cols-2 gap-4">
                    <div>
                      <SearchBar
                        data={[
                          { id: "cloud", title: "Cloud, React.js," },
                          { id: "frontend", title: "Frontend Developer" },
                          { id: "backend", title: "Backend Developer" },
                          { id: "fullstack", title: "Full Stack Developer" },
                          { id: "devops", title: "DevOps Engineer" },
                        ]}
                        placeholder="Cloud, React.js,"
                        onSelect={(item:SearchItem) => console.log("Selected:", item)}
                        maxResults={5}
                        className="w-full"
                      />
                    </div>

                    <div>
                      <Combobox
                        items={experienceOptions}
                        value={selectedExperience}
                        onValueChange={setSelectedExperience}
                        placeholder="Select experience"
                        className="w-full h-12 bg-white border-gray-200 rounded-lg"
                      />
                    </div>
                  </div>

                  <div className="grid grid-cols-2 gap-4">
                    <div>
                      <Combobox
                        items={skillsOptions}
                        value={selectedSkills}
                        onValueChange={setSelectedSkills}
                        placeholder="Select skills"
                        className="w-full h-12 bg-white border-gray-200 rounded-lg"
                      />
                    </div>

                    <div>
                      <SearchBar
                        data={locations}
                        placeholder="Enter location"
                        onSelect={handleLocationSelect}
                        maxResults={5}
                        className="w-full"
                      />
                    </div>
                  </div>

                  <Button
                    onClick={handleSearch}
                    className="w-full h-12 bg-blue-600 hover:bg-blue-700 text-white font-semibold text-base rounded-lg shadow-md hover:shadow-lg transition-all duration-200"
                  >
                    Search
                  </Button>
                </div>
              </div>

              <div className="md:hidden space-y-4">
                <div>
                  <SearchBar
                    data={[
                      { id: "cloud", title: "Cloud, React.js," },
                      { id: "frontend", title: "Frontend Developer" },
                      { id: "backend", title: "Backend Developer" },
                      { id: "fullstack", title: "Full Stack Developer" },
                      { id: "devops", title: "DevOps Engineer" },
                    ]}
                    placeholder="Cloud, React.js,"
                    onSelect={(item:SearchItem) => console.log("Selected:", item)}
                    maxResults={5}
                    className="w-full"
                  />
                </div>

                <div>
                  <Combobox
                    items={experienceOptions}
                    value={selectedExperience}
                    onValueChange={setSelectedExperience}
                    placeholder="Select experience"
                    className="w-full h-12 bg-white border-gray-200 rounded-lg"
                  />
                </div>

                <div>
                  <Combobox
                    items={skillsOptions}
                    value={selectedSkills}
                    onValueChange={setSelectedSkills}
                    placeholder="Select skills"
                    className="w-full h-12 bg-white border-gray-200 rounded-lg"
                  />
                </div>

                <div>
                  <SearchBar
                    data={locations}
                    placeholder="Enter location"
                    onSelect={handleLocationSelect}
                    maxResults={5}
                    className="w-full"
                  />
                </div>

                <Button
                  onClick={handleSearch}
                  className="w-full h-14 bg-blue-600 hover:bg-blue-700 text-white font-semibold text-lg rounded-lg shadow-md hover:shadow-lg transition-all duration-200"
                >
                  Search
                </Button>
              </div>
            </div>
          </div>
        </div>
      </main>
    </div>
  )
}

export default JobZipLanding
