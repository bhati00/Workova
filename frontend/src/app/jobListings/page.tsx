"use client";
import React, { useState, ReactNode } from 'react'
import { ChevronDown, ChevronUp, MapPin, Briefcase, Clock, Filter, X, Building, DollarSign } from 'lucide-react'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card'
import { Combobox } from '@/components/ui/combobox'
import SearchBar from '@/components/ui/searchBar'
import { useSearchParams } from 'next/navigation'

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
    <div className="border-b border-border last:border-b-0">
      <button
        onClick={() => setIsOpen(!isOpen)}
        className="w-full flex items-center justify-between p-4 text-left hover:bg-muted/50 transition-colors"
      >
        <div className="flex items-center gap-3">
          <Icon className="w-5 h-5 text-muted-foreground" />
          <span className="font-medium text-foreground">{title}</span>
        </div>
        {isOpen ? (
          <ChevronUp className="w-5 h-5 text-muted-foreground" />
        ) : (
          <ChevronDown className="w-5 h-5 text-muted-foreground" />
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

// Job Card Component
interface Job {
  id: number;
  title: string;
  company: string;
  location: string;
  type: string;
  workMode: string;
  salary: string;
  posted: string;
  skills: string[];
  description: string;
}

interface JobCardProps {
  job: Job;
}

const JobCard: React.FC<JobCardProps> = ({ job }) => (
  <Card className="hover:shadow-md transition-shadow">
    <CardHeader className="pb-3">
      <div className="flex justify-between items-start">
        <div className="space-y-1">
          <CardTitle className="text-xl leading-tight">{job.title}</CardTitle>
          <div className="flex items-center gap-2 text-muted-foreground">
            <Building className="w-4 h-4" />
            <span className="font-medium">{job.company}</span>
          </div>
        </div>
        <div className="text-sm text-muted-foreground text-right">
          <div className="flex items-center gap-1">
            <Clock className="w-4 h-4" />
            {job.posted}
          </div>
        </div>
      </div>
    </CardHeader>
    
    <CardContent className="space-y-4">
      {/* Job Details */}
      <div className="flex flex-wrap gap-4 text-sm text-muted-foreground">
        <div className="flex items-center gap-1">
          <MapPin className="w-4 h-4" />
          {job.location}
        </div>
        <div className="flex items-center gap-1">
          <Briefcase className="w-4 h-4" />
          {job.type} • {job.workMode}
        </div>
        <div className="flex items-center gap-1">
          <DollarSign className="w-4 h-4" />
          <span className="font-semibold text-foreground">{job.salary}</span>
        </div>
      </div>

      {/* Job Description */}
      <p className="text-sm text-muted-foreground leading-relaxed">
        {job.description}
      </p>

      {/* Skills */}
      <div className="flex flex-wrap gap-2">
        {job.skills.slice(0, 4).map((skill, index) => (
          <span key={index} className="px-2 py-1 bg-secondary text-secondary-foreground text-xs rounded-md font-medium">
            {skill}
          </span>
        ))}
        {job.skills.length > 4 && (
          <span className="px-2 py-1 bg-muted text-muted-foreground text-xs rounded-md">
            +{job.skills.length - 4} more
          </span>
        )}
      </div>

      {/* Actions */}
      <div className="flex gap-3 pt-2">
        <Button variant="primary-blue" size="sm" className="flex-1">
          Apply Now
        </Button>
        <Button variant="outline-blue" size="sm" className="flex-1">
          Save Job
        </Button>
      </div>
    </CardContent>
  </Card>
)

// Page Header Component
const PageHeader: React.FC = () => {
  const searchParams = useSearchParams()
  
  const category = searchParams.get('category') || ''
  const skills = searchParams.get('skills')?.split(',') || []
  const location = searchParams.get('location') || ''
  
  const generateDescription = () => {
    const parts = []
    
    if (category) parts.push(`${category} jobs`)
    if (skills.length > 0) {
      if (skills.length === 1) {
        parts.push(`requiring ${skills[0]} skills`)
      } else if (skills.length === 2) {
        parts.push(`requiring ${skills[0]} and ${skills[1]} skills`)
      } else {
        parts.push(`requiring ${skills[0]}, ${skills[1]}, and ${skills.length - 2} other skills`)
      }
    }
    if (location) parts.push(`in ${location}`)
    
    if (parts.length === 0) {
      return "Discover your next career opportunity from our curated job listings"
    }
    
    return `Find ${parts.join(' ')}`
  }

  return (
    <Card>
      <CardHeader>
        <CardTitle className="text-3xl font-bold">🚀 Your Gateway to the Most Relevant Jobs</CardTitle>
        <CardDescription className="text-lg">
          {generateDescription()}
        </CardDescription>
      </CardHeader>
    </Card>
  )
}

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

  // Enhanced mock job data
  const mockJobs: Job[] = [
    {
      id: 1,
      title: "Senior Software Engineer",
      company: "Tech Corp",
      location: "San Francisco, CA",
      type: "Full Time",
      workMode: "Hybrid",
      salary: "$120k - $180k",
      posted: "2 days ago",
      skills: ["React", "Node.js", "TypeScript", "AWS", "Docker"],
      description: "Join our engineering team to build scalable web applications using modern technologies. You'll work on challenging projects that impact millions of users worldwide."
    },
    {
      id: 2,
      title: "Frontend Developer",
      company: "StartupXYZ",
      location: "Remote",
      type: "Contract",
      workMode: "Remote",
      salary: "$80k - $120k",
      posted: "1 week ago",
      skills: ["Vue.js", "JavaScript", "CSS", "Git", "Figma"],
      description: "Looking for a creative frontend developer to help us build beautiful, responsive user interfaces. Experience with modern JavaScript frameworks required."
    },
    {
      id: 3,
      title: "DevOps Engineer",
      company: "CloudTech Solutions",
      location: "New York, NY",
      type: "Full Time",
      workMode: "Onsite",
      salary: "$100k - $150k",
      posted: "3 days ago",
      skills: ["Docker", "Kubernetes", "AWS", "Python", "Terraform"],
      description: "Help us scale our infrastructure and improve deployment processes. You'll work with cutting-edge cloud technologies and automation tools."
    },
    {
      id: 4,
      title: "Product Manager",
      company: "InnovateCorp",
      location: "Austin, TX",
      type: "Full Time",
      workMode: "Hybrid",
      salary: "$90k - $130k",
      posted: "5 days ago",
      skills: ["Product Strategy", "Analytics", "Agile", "User Research"],
      description: "Drive product vision and strategy for our flagship products. Work closely with engineering, design, and business teams to deliver exceptional user experiences."
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
    <div className="min-h-screen bg-background">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-6">
        {/* Page Header */}
        <div className="mb-6">
          <PageHeader />
        </div>

        {/* Main Content */}
        <div className="grid grid-cols-1 lg:grid-cols-4 gap-6">
          {/* Filters Sidebar */}
          <div className="lg:col-span-1">
            <Card className="sticky top-6">
              <CardHeader className="pb-3">
                <div className="flex items-center gap-2">
                  <Filter className="w-5 h-5 text-muted-foreground" />
                  <CardTitle className="text-lg">Filters</CardTitle>
                </div>
              </CardHeader>
              
              <CardContent className="p-0">
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
                            className="inline-flex items-center gap-1 px-2 py-1 bg-secondary text-secondary-foreground text-xs rounded-md"
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
                            className="w-4 h-4 text-blue-600 border-border rounded focus:ring-blue-500"
                          />
                          <span className="text-sm text-foreground">{type.label}</span>
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
                            className="w-4 h-4 text-blue-600 border-border rounded focus:ring-blue-500"
                          />
                          <span className="text-sm text-foreground">{level.label}</span>
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
                        <span className="inline-flex items-center gap-1 px-2 py-1 bg-secondary text-secondary-foreground text-xs rounded-md">
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
                            className="w-4 h-4 text-blue-600 border-border focus:ring-blue-500"
                          />
                          <span className="text-sm text-foreground">{duration.label}</span>
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
                          className="w-4 h-4 text-blue-600 border-border rounded focus:ring-blue-500"
                        />
                        <span className="text-sm text-foreground">Remote Only</span>
                      </label>
                      <label className="flex items-center gap-2">
                        <input
                          type="checkbox"
                          checked={filters.visaSponsorship}
                          onChange={(e) => handleFilterChange('visaSponsorship', e.target.checked)}
                          className="w-4 h-4 text-blue-600 border-border rounded focus:ring-blue-500"
                        />
                        <span className="text-sm text-foreground">Visa Sponsorship</span>
                      </label>
                    </div>
                  </FilterAccordion>
                </div>

                {/* Filter Actions */}
                <div className="p-4 border-t border-border space-y-2">
                  <Button 
                    onClick={applyFilters}
                    variant="primary-blue"
                    className="w-full"
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
              </CardContent>
            </Card>
          </div>

          {/* Job Listings */}
          <div className="lg:col-span-3">
            <div className="mb-6">
              <div className="flex justify-between items-center">
                <p className="text-muted-foreground">
                  Showing <span className="font-semibold text-foreground">{mockJobs.length}</span> jobs
                  {hasFilters && ' (filtered)'}
                </p>
                <Button variant="outline-blue" size="sm">
                  <Filter className="w-4 h-4" />
                  Sort by: Relevance
                </Button>
              </div>
            </div>
            
            <div className="space-y-4">
              {mockJobs.map((job) => (
                <JobCard key={job.id} job={job} />
              ))}
            </div>

            {/* Load More */}
            <div className="mt-8 text-center">
              <Button variant="outline" size="lg" className="px-8">
                Load More Jobs
              </Button>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}