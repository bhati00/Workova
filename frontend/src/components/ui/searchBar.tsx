"use client"

import * as React from "react"
import { Command, CommandEmpty, CommandGroup, CommandInput, CommandItem, CommandList } from "@/components/ui/command"

// Types
interface SearchItem {
  id: string | number
  title: string
  subtitle?: string
}

interface SearchBarProps {
  data: SearchItem[]
  placeholder?: string
  onSelect: (item: SearchItem) => void
  maxResults?: number
  className?: string
  emptyText?: string
}

const SearchBar: React.FC<SearchBarProps> = ({
  data,
  placeholder = "Search...",
  onSelect,
  maxResults = 10,
  className = "",
  emptyText = "No results found.",
}) => {
  const [value, setValue] = React.useState("")
  const [isOpen, setIsOpen] = React.useState(false)
  const commandRef = React.useRef<HTMLDivElement>(null)

  const handleSelect = (currentValue: string) => {
    const selectedItem = data.find((item) => item.id.toString() === currentValue)

    if (selectedItem) {
      setValue(selectedItem.title)
      onSelect(selectedItem)
      setIsOpen(false)
    }
  }

  const handleInputFocus = () => {
    setIsOpen(true)
  }

  const handleInputChange = (newValue: string) => {
    setValue(newValue)
    setIsOpen(newValue.length > 0 || data.length > 0)
  }

  React.useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (commandRef.current && !commandRef.current.contains(event.target as Node)) {
        setIsOpen(false)
      }
    }

    document.addEventListener("mousedown", handleClickOutside)
    return () => {
      document.removeEventListener("mousedown", handleClickOutside)
    }
  }, [])

  // Filter and limit results
  const filteredData = React.useMemo(() => {
    if (!value.trim()) return data.slice(0, maxResults)

    return data
      .filter(
        (item) =>
          item.title.toLowerCase().includes(value.toLowerCase()) ||
          (item.subtitle && item.subtitle.toLowerCase().includes(value.toLowerCase())),
      )
      .slice(0, maxResults)
  }, [data, value, maxResults])

  return (
    <div ref={commandRef} className="relative">
      <Command className={`${className}`}>
        <CommandInput
          placeholder={placeholder}
          value={value}
          onValueChange={handleInputChange}
          onFocus={handleInputFocus}
          className="h-12 px-3 border border-gray-200 rounded-lg bg-white text-sm placeholder:text-gray-500"
        />
        {isOpen && (
          <CommandList className="absolute top-full left-0 right-0 z-50 bg-white border border-gray-200 rounded-lg shadow-lg mt-1 max-h-60 overflow-auto">
            <CommandEmpty className="py-6 text-center text-sm text-gray-500">{emptyText}</CommandEmpty>
            <CommandGroup>
              {filteredData.map((item) => (
                <CommandItem
                  key={item.id}
                  value={item.id.toString()}
                  onSelect={handleSelect}
                  className="cursor-pointer px-3 py-2 hover:bg-gray-50 data-[selected=true]:bg-blue-50 data-[selected=true]:text-blue-900"
                >
                  <div className="flex flex-col">
                    <span className="text-sm font-medium">{item.title}</span>
                    {item.subtitle && <span className="text-xs text-gray-500">{item.subtitle}</span>}
                  </div>
                </CommandItem>
              ))}
            </CommandGroup>
          </CommandList>
        )}
      </Command>
    </div>
  )
}

export default SearchBar
export type { SearchItem }
