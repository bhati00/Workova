"use client";
import Link from "next/link";
import { useState } from "react";
import Select, { MultiValue, ActionMeta } from "react-select";
import { Button } from "@/components/ui/button";
import { Combobox } from "@/components/ui/combobox";
import SearchBar, { type SearchItem } from "@/components/ui/searchBar";
import  { useRouter } from "next/navigation";
import { Search } from "lucide-react";


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
];

const experienceOptions = [
  { value: "all", label: "Select experience" },
  { value: "fresher", label: "Fresher (less than 1 year)" },
  { value: "1year", label: "1 year" },
  { value: "2years", label: "2 years" },
  { value: "3years", label: "3 years" },
  { value: "4years", label: "4 years" },
  { value: "5years", label: "5 years" },
];

const skillsOptions = [
  { value: "react", label: "React.js" },
  { value: "nodejs", label: "Node.js" },
  { value: "python", label: "Python" },
  { value: "javascript", label: "JavaScript" },
  { value: "typescript", label: "TypeScript" },
  { value: "aws", label: "AWS" },
  { value: "docker", label: "Docker" },
  { value: "kubernetes", label: "Kubernetes" },
  { value: "mysql", label: "MySQL" },
  { value: "postgresql", label: "PostgreSQL" },
  { value: "mongodb", label: "MongoDB" },
  { value: "redis", label: "Redis" },
  { value: "graphql", label: "GraphQL" },
  { value: "nextjs", label: "Next.js" },
  { value: "angular", label: "Angular" },
  { value: "vue", label: "Vue.js" },
];

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
];

interface SkillOption {
  value: string;
  label: string;
}

// Custom styles for react-select to match the existing design
const customSelectStyles = {
  control: (provided: Record<string, unknown>, state: { isFocused: boolean }) => ({
    ...provided,
    minHeight: '48px',
    height: '48px',
    border: state.isFocused ? '1px solid #3b82f6' : '1px solid #e5e7eb',
    borderRadius: '8px',
    backgroundColor: '#ffffff',
    boxShadow: state.isFocused ? '0 0 0 3px rgba(59, 130, 246, 0.1)' : 'none',
    '&:hover': {
      border: state.isFocused ? '1px solid #3b82f6' : '1px solid #d1d5db',
    },
    transition: 'all 0.2s ease-in-out',
  }),
  valueContainer: (provided: Record<string, unknown>) => ({
    ...provided,
    padding: '2px 8px',
    height: '44px',
  }),
  input: (provided: Record<string, unknown>) => ({
    ...provided,
    margin: '0px',
    padding: '0px',
    color: '#374151',
  }),
  indicatorSeparator: () => ({
    display: 'none',
  }),
  indicatorsContainer: (provided: Record<string, unknown>) => ({
    ...provided,
    height: '44px',
  }),
  menu: (provided: Record<string, unknown>) => ({
    ...provided,
    backgroundColor: '#ffffff',
    border: '1px solid #e5e7eb',
    borderRadius: '8px',
    boxShadow: '0 10px 15px -3px rgba(0, 0, 0, 0.1), 0 4px 6px -2px rgba(0, 0, 0, 0.05)',
    zIndex: 9999,
  }),
  option: (provided: Record<string, unknown>, state: { isSelected: boolean; isFocused: boolean }) => ({
    ...provided,
    backgroundColor: state.isSelected 
      ? '#dbeafe' 
      : state.isFocused 
      ? '#f9fafb' 
      : '#ffffff',
    color: state.isSelected ? '#1e40af' : '#374151',
    padding: '8px 12px',
    cursor: 'pointer',
    '&:hover': {
      backgroundColor: '#f9fafb',
    },
  }),
  multiValue: (provided: Record<string, unknown>) => ({
    ...provided,
    backgroundColor: '#e0e7ff',
    borderRadius: '6px',
    margin: '2px',
  }),
  multiValueLabel: (provided: Record<string, unknown>) => ({
    ...provided,
    color: '#3730a3',
    fontSize: '14px',
    padding: '2px 6px',
  }),
  multiValueRemove: (provided: Record<string, unknown>) => ({
    ...provided,
    color: '#6366f1',
    '&:hover': {
      backgroundColor: '#c7d2fe',
      color: '#4338ca',
    },
  }),
  placeholder: (provided: Record<string, unknown>) => ({
    ...provided,
    color: '#9ca3af',
    fontSize: '14px',
  }),
};

const JobZipLanding = () => {
  const [selectedCategory, setSelectedCategory] = useState("");
  const [selectedExperience, setSelectedExperience] = useState("");
  const [selectedSkills, setSelectedSkills] = useState<SkillOption[]>([]);

  const handleLocationSelect = (location: SearchItem) => {
    console.log("Selected location:", location);
  };

  const handleSkillsChange = (selectedOptions: MultiValue<SkillOption>, actionMeta: ActionMeta<SkillOption>) => {
    setSelectedSkills([...selectedOptions]);
  };

const router = useRouter()
  const handleSearch = () => {
     const searchParams = new URLSearchParams({
    category: selectedCategory,
    experience: selectedExperience,
    skills: selectedSkills.map((skill) => skill.value).join(','),
  });
    router.push(`/jobListings?${searchParams.toString()}`);
  };

  return (
      <main className="container mx-auto flex min-h-screen flex-col items-center">
        <div className="max-w-6xl px-6 pb-12 pt-8 sm:pb-40 text-center relative w-full">
          <h1 className="mt-24 sm:mt-16 lg:mt-8 heading-1 animate-fade-in-up">
            Find your dream job now
          </h1>

          <p
            className="paragraph paragraph-center animate-fade-in-up"
            style={{ animationDelay: "0.2s" }}
          >
            5 lakh+ jobs for you to explore
          </p>

          <div
            className="mt-12 animate-fade-in-up"
            style={{ animationDelay: "0.4s" }}
          >
            <div className="bg-white rounded-2xl shadow-xl border border-gray-100 p-6 max-w-6xl mx-auto">
            {/*Desktop Menu*/}
              <div className="hidden lg:block">
                <div className="grid grid-cols-5 gap-4 items-center">
                  <div>
                    <SearchBar
                      data={locations}
                      placeholder="Enter location"
                      onSelect={handleLocationSelect}
                      maxResults={5}
                      className="w-full border border-gray-200 rounded-lg p-[7px]"
                    />
                  </div>
                  <div>
                    <Combobox
                      items={categories}
                      value={selectedCategory}
                      onValueChange={setSelectedCategory}
                      placeholder="Select Category"
                      className="w-full h-12 bg-white border-gray-200 rounded-lg"
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
                    <Select
                      isMulti
                      options={skillsOptions}
                      value={selectedSkills}
                      onChange={handleSkillsChange}
                      placeholder="Select skills"
                      styles={customSelectStyles}
                      className="w-full"
                      classNamePrefix="select"
                      isClearable
                      isSearchable
                      closeMenuOnSelect={false}
                    />
                  </div>

                  <div>
                    <Button
                      onClick={handleSearch}
                     variant="primary-blue"
                    >
                      <Search className="w-4 h-4" />
                      Search
                    </Button>
                  </div>
                </div>
              </div>
            {/*Tablet Menu*/}
              <div className="hidden md:block lg:hidden">
                <div className="space-y-4">
                  <div className="grid grid-cols-2 gap-4">
                    <div>
                      <SearchBar
                        data={locations}
                        placeholder="Enter location"
                        onSelect={handleLocationSelect}
                        maxResults={5}
                        className="w-full border border-gray-200 rounded-lg p-[7px]"
                      />
                    </div>

                    <div>
                      <Combobox
                        items={categories}
                        value={selectedCategory}
                        onValueChange={setSelectedCategory}
                        placeholder="Select Category"
                        className="w-full h-12 bg-white border-gray-200 rounded-lg"
                      />
                    </div>
                  </div>

                  <div className="grid grid-cols-2 gap-4">
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
                      <Select
                        isMulti
                        options={skillsOptions}
                        value={selectedSkills}
                        onChange={handleSkillsChange}
                        placeholder="Select skills"
                        styles={customSelectStyles}
                        className="w-full"
                        classNamePrefix="select"
                        isClearable
                        isSearchable
                        closeMenuOnSelect={false}
                      />
                    </div>
                  </div>

                  <Button
                    onClick={handleSearch}
                   variant="primary-blue"
                  >
                    <Search className="w-4 h-4" />
                    Search
                  </Button>
                </div>
              </div>
              {/* {Mobile Menu} */}
              <div className="md:hidden space-y-4">
                <div>
                  <SearchBar
                    data={locations}
                    placeholder="Enter location"
                    onSelect={handleLocationSelect}
                    maxResults={5}
                    className="w-full border border-gray-200 rounded-lg p-[7px]"
                  />
                </div>

                <div>
                  <Combobox
                    items={categories}
                    value={selectedCategory}
                    onValueChange={setSelectedCategory}
                    placeholder="Select Category"
                    className="w-full h-12 bg-white border-gray-200 rounded-lg"
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
                  <Select
                    isMulti
                    options={skillsOptions}
                    value={selectedSkills}
                    onChange={handleSkillsChange}
                    placeholder="Select skills"
                    styles={customSelectStyles}
                    className="w-full"
                    classNamePrefix="select"
                    isClearable
                    isSearchable
                    closeMenuOnSelect={false}
                  />
                </div>

                <Button
                  onClick={handleSearch}
                  variant="primary-blue"
                >
                  <Search className="w-4 h-4" />
                  Search
                </Button>
              </div>
            </div>
          </div>
        </div>
      </main>
  );
};

export default JobZipLanding;