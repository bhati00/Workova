"use client";
import React, { useState, ReactNode } from 'react'
import { ChevronDown, ChevronUp, MapPin, Briefcase, Clock, Filter, X } from 'lucide-react'
import { Button } from '@/components/ui/button'
import { Combobox } from '@/components/ui/combobox'
import SearchBar from '@/components/ui/searchBar'
import Select from 'react-select'

// Constants from backend
const WORK_MODES = [
  { value: '1', label: 'Remote' },
  { value: '2', label: 'Onsite' },
  { value: '3', label: 'Hybrid' }
]

const JOB_TYPES = [
  { value: '1', label: 'Full Time' },
  { value: '2', label: 'Part Time' },
  { value: '3', label: 'Contract' },
  { value: '4', label: 'Internship' },
  { value: '5', label: 'Temporary' }
]

const EXPERIENCE_LEVELS = [
  { value: '1', label: 'Entry Level' },
  { value: '2', label: 'Mid Level' },
  { value: '3', label: 'Senior Level' },
  { value: '4', label: 'Lead' },
  { value: '5', label: 'Executive' }
]

const POSTED_DURATION = [
  { value: '1', label: 'Last 24 hours' },
  { value: '7', label: 'Last 7 days' },
  { value: '30', label: 'Last 30 days' },
  { value: '90', label: 'Last 3 months' }
]

// Mock data for skills and locations
const SKILLS_DATA = [
  { id: 1, title: 'JavaScript', subtitle: 'Programming Language' },
  { id: 2, title: 'React', subtitle: 'Frontend Framework' },
  { id: 3, title: 'Node.js', subtitle: 'Backend Runtime' },
  { id: 4, title: 'Python', subtitle: 'Programming Language' },
  { id: 5, title: 'Go', subtitle: 'Programming Language' },
  { id: 6, title: 'Docker', subtitle: 'Containerization' },
  { id: 7, title: 'Kubernetes', subtitle: 'Orchestration' },
  { id: 8, title: 'AWS', subtitle: 'Cloud Platform' }
]

const LOCATION_DATA = [
  { id: 1, title: 'San Francisco, CA', subtitle: 'United States' },
  { id: 2, title: 'New York, NY', subtitle: 'United States' },
  { id: 3, title: 'London', subtitle: 'United Kingdom' },
  { id: 4, title: 'Berlin', subtitle: 'Germany' },
  { id: 5, title: 'Toronto', subtitle: 'Canada' },
  { id: 6, title: 'Sydney', subtitle: 'Australia' }
]

// Filter types
interface FilterState {
  workMode: string;
  skills: string[];
  jobType: string[];
  experienceLevel: string[];
  visaSponsorship: boolean;
  isRemote: boolean;
  location: string;
  postedDuration: string;
}

type FilterValue = string | string[] | boolean;

// Accordion Component
interface FilterAccordionProps {
  title: string;
  icon: React.ComponentType<{ className?: string }>;
  children: ReactNode;
  defaultOpen?: boolean;
}

const FilterAccordion: React.FC<FilterAccordionProps> = ({ title, icon: Icon, children, defaultOpen = false }) => {
  const [isOpen, setIsOpen] = useState(defaultOpen)
  
  return (
    <div className="border-b border-gray-200 last:border-b-0">
      <button
        onClick={() => setIsOpen(!isOpen)}
        className="w-full flex items-center justify-between p-4 text-left hover:bg-gray-50 transition-colors"
      >
        <div className="flex items-center gap-3">
          <Icon className="w-5 h-5 text-gray-600" />
          <span className="font-medium text-gray-900">{title}</span>
        </div>
        {isOpen ? (
          <ChevronUp className="w-5 h-5 text-gray-500" />
        ) : (
          <ChevronDown className="w-5 h-5 text-gray-500" />
        )}
      </button>
      {isOpen && (
        <div className="px-4 pb-4">
          {children}
        </div>
      )}
    </div>
  )
}

// Multi-Select Component for Skills


// Job Card Component (placeholder)
interface Job {
  id: number;
  title: string;
  company: string;
  location: string;
  type: string;
  salary: string;
  posted: string;
  skills: string[];
}

interface JobCardProps {
  job: Job;
}

const JobCard: React.FC<JobCardProps> = ({ job }) => (
  <div className="bg-white rounded-lg border border-gray-200 p-6 hover:shadow-md transition-shadow">
    <div className="flex justify-between items-start mb-3">
      <h3 className="text-lg font-semibold text-gray-900">{job.title}</h3>
      <span className="text-sm text-gray-500">{job.posted}</span>
    </div>
    <p className="text-gray-700 font-medium mb-2">{job.company}</p>
    <div className="flex items-center gap-4 text-sm text-gray-600 mb-3">
      <span className="flex items-center gap-1">
        <MapPin className="w-4 h-4" />
        {job.location}
      </span>
      <span className="flex items-center gap-1">
        <Briefcase className="w-4 h-4" />
        {job.type}
      </span>
    </div>
    <div className="flex flex-wrap gap-2 mb-4">
      {job.skills.slice(0, 3).map((skill, index) => (
        <span key={index} className="px-2 py-1 bg-blue-100 text-blue-800 text-xs rounded-md">
          {skill}
        </span>
      ))}
    </div>
    <div className="flex justify-between items-center">
      <span className="text-lg font-semibold text-gray-900">{job.salary}</span>
      <Button size="sm" className="bg-orange-500 hover:bg-orange-600">
        Apply Now
      </Button>
    </div>
  </div>
)

// Main Component
export default function JobListingsPage() {
  const [filters, setFilters] = useState<FilterState>({
    workMode: '',
    skills: [],
    jobType: [],
    experienceLevel: [],
    visaSponsorship: false,
    isRemote: false,
    location: '',
    postedDuration: ''
  })

  const [hasFilters, setHasFilters] = useState(false)

  // Mock job data
  const mockJobs = [
    {
      id: 1,
      title: "Senior Software Engineer",
      company: "Tech Corp",
      location: "San Francisco, CA",
      type: "Full Time",
      salary: "$120k - $180k",
      posted: "2 days ago",
      skills: ["React", "Node.js", "TypeScript", "AWS"]
    },
    {
      id: 2,
      title: "Frontend Developer",
      company: "StartupXYZ",
      location: "Remote",
      type: "Contract",
      salary: "$80k - $120k",
      posted: "1 week ago",
      skills: ["Vue.js", "JavaScript", "CSS", "Git"]
    },
    {
      id: 3,
      title: "DevOps Engineer",
      company: "CloudTech",
      location: "New York, NY",
      type: "Full Time",
      salary: "$100k - $150k",
      posted: "3 days ago",
      skills: ["Docker", "Kubernetes", "AWS", "Python"]
    }
  ]

  const handleFilterChange = (filterType: keyof FilterState, value: FilterValue) => {
    setFilters(prev => ({
      ...prev,
      [filterType]: value
    }))
    
    // Check if any filters are applied
    const newFilters = { ...filters, [filterType]: value }
    const hasAnyFilter = Object.values(newFilters).some(val => 
      Array.isArray(val) ? val.length > 0 : val !== '' && val !== false
    )
    setHasFilters(hasAnyFilter)
  }

  const clearAllFilters = () => {
    setFilters({
      workMode: '',
      skills: [],
      jobType: [],
      experienceLevel: [],
      visaSponsorship: false,
      isRemote: false,
      location: '',
      postedDuration: ''
    })
    setHasFilters(false)
  }

  const applyFilters = () => {
    // Implementation for applying filters
    console.log('Applying filters:', filters)
  }

  return (
    <div className="min-h-screen bg-gray-50">
      {/* Header */}
      <div className="bg-white border-b border-gray-200">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4">
          <h1 className="text-2xl font-bold text-gray-900">Job Listings</h1>
          <p className="text-gray-600 mt-1">Find your next opportunity</p>
        </div>
      </div>

      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-6">
        <div className="grid grid-cols-1 lg:grid-cols-4 gap-6">
          {/* Filters Sidebar */}
          <div className="lg:col-span-1">
            <div className="bg-white rounded-lg shadow-sm border border-gray-200 sticky top-6">
              <div className="p-4 border-b border-gray-200">
                <div className="flex items-center gap-2">
                  <Filter className="w-5 h-5 text-gray-600" />
                  <h2 className="text-lg font-semibold text-gray-900">Filters</h2>
                </div>
              </div>

              <div className="max-h-[70vh] overflow-y-auto">
                {/* Work Mode Filter */}
                <FilterAccordion title="Work Mode" icon={Briefcase}>
                  <Combobox
                    items={WORK_MODES}
                    value={filters.workMode}
                    onValueChange={(value) => handleFilterChange('workMode', value)}
                    placeholder="Select work mode"
                    className="w-full"
                  />
                </FilterAccordion>

                {/* Skills Filter */}
                <FilterAccordion title="Skills" icon={Briefcase}>
                  <SearchBar
                    data={SKILLS_DATA}
                    placeholder="Search skills..."
                    onSelect={(item) => {
                      const newSkills = [...filters.skills, item.title]
                      handleFilterChange('skills', newSkills)
                    }}
                    maxResults={8}
                    emptyText="No skills found"
                  />
                  {filters.skills.length > 0 && (
                    <div className="flex flex-wrap gap-2 mt-3">
                      {filters.skills.map((skill, index) => (
                        <span
                          key={index}
                          className="inline-flex items-center gap-1 px-2 py-1 bg-blue-100 text-blue-800 text-xs rounded-md"
                        >
                          {skill}
                          <button
                            onClick={() => {
                              const newSkills = filters.skills.filter(s => s !== skill)
                              handleFilterChange('skills', newSkills)
                            }}
                          >
                            <X className="w-3 h-3" />
                          </button>
                        </span>
                      ))}
                    </div>
                  )}
                </FilterAccordion>

                {/* Job Type Filter */}
                <FilterAccordion title="Job Type" icon={Briefcase}>
                  <div className="space-y-2">
                    {JOB_TYPES.map((type) => (
                      <label key={type.value} className="flex items-center gap-2">
                        <input
                          type="checkbox"
                          checked={filters.jobType.includes(type.value)}
                          onChange={(e) => {
                            const newJobTypes = e.target.checked
                              ? [...filters.jobType, type.value]
                              : filters.jobType.filter(jt => jt !== type.value)
                            handleFilterChange('jobType', newJobTypes)
                          }}
                          className="w-4 h-4 text-orange-500 border-gray-300 rounded focus:ring-orange-500"
                        />
                        <span className="text-sm text-gray-700">{type.label}</span>
                      </label>
                    ))}
                  </div>
                </FilterAccordion>

                {/* Experience Level Filter */}
                <FilterAccordion title="Experience Level" icon={Briefcase}>
                  <div className="space-y-2">
                    {EXPERIENCE_LEVELS.map((level) => (
                      <label key={level.value} className="flex items-center gap-2">
                        <input
                          type="checkbox"
                          checked={filters.experienceLevel.includes(level.value)}
                          onChange={(e) => {
                            const newLevels = e.target.checked
                              ? [...filters.experienceLevel, level.value]
                              : filters.experienceLevel.filter(el => el !== level.value)
                            handleFilterChange('experienceLevel', newLevels)
                          }}
                          className="w-4 h-4 text-orange-500 border-gray-300 rounded focus:ring-orange-500"
                        />
                        <span className="text-sm text-gray-700">{level.label}</span>
                      </label>
                    ))}
                  </div>
                </FilterAccordion>

                {/* Location Filter */}
                <FilterAccordion title="Location" icon={MapPin}>
                  <SearchBar
                    data={LOCATION_DATA}
                    placeholder="Search location..."
                    onSelect={(item) => handleFilterChange('location', item.title)}
                    maxResults={6}
                    emptyText="No locations found"
                  />
                  {filters.location && (
                    <div className="mt-3">
                      <span className="inline-flex items-center gap-1 px-2 py-1 bg-green-100 text-green-800 text-xs rounded-md">
                        {filters.location}
                        <button onClick={() => handleFilterChange('location', '')}>
                          <X className="w-3 h-3" />
                        </button>
                      </span>
                    </div>
                  )}
                </FilterAccordion>

                {/* Posted Duration Filter */}
                <FilterAccordion title="Posted" icon={Clock}>
                  <div className="space-y-2">
                    {POSTED_DURATION.map((duration) => (
                      <label key={duration.value} className="flex items-center gap-2">
                        <input
                          type="radio"
                          name="postedDuration"
                          value={duration.value}
                          checked={filters.postedDuration === duration.value}
                          onChange={(e) => handleFilterChange('postedDuration', e.target.value)}
                          className="w-4 h-4 text-orange-500 border-gray-300 focus:ring-orange-500"
                        />
                        <span className="text-sm text-gray-700">{duration.label}</span>
                      </label>
                    ))}
                  </div>
                </FilterAccordion>

                {/* Boolean Filters */}
                <FilterAccordion title="Other Filters" icon={Filter}>
                  <div className="space-y-3">
                    <label className="flex items-center gap-2">
                      <input
                        type="checkbox"
                        checked={filters.isRemote}
                        onChange={(e) => handleFilterChange('isRemote', e.target.checked)}
                        className="w-4 h-4 text-orange-500 border-gray-300 rounded focus:ring-orange-500"
                      />
                      <span className="text-sm text-gray-700">Remote Only</span>
                    </label>
                    <label className="flex items-center gap-2">
                      <input
                        type="checkbox"
                        checked={filters.visaSponsorship}
                        onChange={(e) => handleFilterChange('visaSponsorship', e.target.checked)}
                        className="w-4 h-4 text-orange-500 border-gray-300 rounded focus:ring-orange-500"
                      />
                      <span className="text-sm text-gray-700">Visa Sponsorship</span>
                    </label>
                  </div>
                </FilterAccordion>
              </div>

              {/* Filter Actions */}
              <div className="p-4 border-t border-gray-200 space-y-2">
                <Button 
                  onClick={applyFilters}
                  className="w-full bg-orange-500 hover:bg-orange-600"
                >
                  Apply Filters
                </Button>
                {hasFilters && (
                  <Button 
                    onClick={clearAllFilters}
                    variant="outline"
                    className="w-full"
                  >
                    Clear All Filters
                  </Button>
                )}
              </div>
            </div>
          </div>

          {/* Job Listings */}
          <div className="lg:col-span-3">
            <div className="mb-4">
              <p className="text-gray-600">
                Showing {mockJobs.length} jobs
                {hasFilters && ' (filtered)'}
              </p>
            </div>
            
            <div className="grid gap-6">
              {mockJobs.map((job) => (
                <JobCard key={job.id} job={job} />
              ))}
            </div>

            {/* Load More */}
            <div className="mt-8 text-center">
              <Button variant="outline" className="px-8">
                Load More Jobs
              </Button>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}