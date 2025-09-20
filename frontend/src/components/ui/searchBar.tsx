"use client"

import * as React from "react"
import { Command, CommandEmpty, CommandGroup, CommandInput, CommandItem, CommandList } from "@/components/ui/command"

// Types
interface SearchItem {
  id: string | number;
  title: string;
  subtitle?: string;
}

interface SearchBarProps {
  data: SearchItem[];
  placeholder?: string;
  onSelect: (item: SearchItem) => void;
  maxResults?: number;
  className?: string;
  emptyText?: string;
}

const SearchBar: React.FC<SearchBarProps> = ({
  data,
  placeholder = "Search...",
  onSelect,
  maxResults = 10,
  className = "",
  emptyText = "No results found."
}) => {
  const [value, setValue] = React.useState("")

  const handleSelect = (currentValue: string) => {
    const selectedItem = data.find(item => 
      item.id.toString() === currentValue
    );
    
    if (selectedItem) {
      setValue(selectedItem.title);
      onSelect(selectedItem);
    }
  };

  // Filter and limit results
  const filteredData = React.useMemo(() => {
    if (!value.trim()) return data.slice(0, maxResults);
    
    return data
      .filter(item => 
        item.title.toLowerCase().includes(value.toLowerCase()) ||
        (item.subtitle && item.subtitle.toLowerCase().includes(value.toLowerCase()))
      )
      .slice(0, maxResults);
  }, [data, value, maxResults]);

  return (
    <Command className={`search-container ${className}`}>
      <CommandInput 
        placeholder={placeholder}
        value={value}
        onValueChange={setValue}
      />
      <CommandList>
        <CommandEmpty>{emptyText}</CommandEmpty>
        <CommandGroup>
          {filteredData.map((item) => (
            <CommandItem
              key={item.id}
              value={item.id.toString()}
              onSelect={handleSelect}
              className="cursor-pointer"
            >
              <div className="flex flex-col">
                <span className="text-sm font-medium">{item.title}</span>
                {item.subtitle && (
                  <span className="text-xs text-muted-foreground">{item.subtitle}</span>
                )}
              </div>
            </CommandItem>
          ))}
        </CommandGroup>
      </CommandList>
    </Command>
  )
}

export default SearchBar;