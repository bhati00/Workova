import * as React from "react"
import { Slot } from "@radix-ui/react-slot"
import { cva, type VariantProps } from "class-variance-authority"

import { cn } from "@/lib/utils"

const buttonVariants = cva(
  "inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-md text-sm font-medium cursor-pointer transition-all disabled:pointer-events-none disabled:opacity-50 disabled:cursor-not-allowed [&_svg]:pointer-events-none [&_svg:not([class*='size-'])]:size-4 shrink-0 [&_svg]:shrink-0 outline-none focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] aria-invalid:ring-destructive/20 dark:aria-invalid:ring-destructive/40 aria-invalid:border-destructive",
  {
    variants: {
      variant: {
        // Original shadcn variants (kept for compatibility)
        default: "bg-primary text-primary-foreground hover:bg-primary/90",
        destructive:
          "bg-destructive text-white hover:bg-destructive/90 focus-visible:ring-destructive/20 dark:focus-visible:ring-destructive/40 dark:bg-destructive/60",
        outline:
          "border bg-background shadow-xs hover:bg-accent hover:text-accent-foreground dark:bg-input/30 dark:border-input dark:hover:bg-input/50",
        secondary:
          "bg-secondary text-secondary-foreground hover:bg-secondary/80",
        ghost:
          "hover:bg-accent hover:text-accent-foreground dark:hover:bg-accent/50",
        link: "text-primary underline-offset-4 hover:underline",
        
        // New Naukri-inspired variants
        "primary-blue": 
          "bg-blue-600 text-white hover:bg-blue-700 active:bg-blue-800 focus-visible:ring-blue-600/30 shadow-sm font-semibold dark:bg-blue-500 dark:hover:bg-blue-600",
        "primary-orange": 
          "bg-orange-500 text-white hover:bg-orange-600 active:bg-orange-700 focus-visible:ring-orange-500/30 shadow-sm font-semibold dark:bg-orange-500 dark:hover:bg-orange-600",
        "outline-blue": 
          "border-2 border-blue-600 text-blue-600 bg-transparent hover:bg-blue-50 active:bg-blue-100 focus-visible:ring-blue-600/30 font-semibold dark:border-blue-400 dark:text-blue-400 dark:hover:bg-blue-950",
        
        // Additional professional variants
        "outline-orange": 
          "border-2 border-orange-500 text-orange-500 bg-transparent hover:bg-orange-50 active:bg-orange-100 focus-visible:ring-orange-500/30 font-semibold dark:border-orange-400 dark:text-orange-400 dark:hover:bg-orange-950",
        "success": 
          "bg-green-600 text-white hover:bg-green-700 active:bg-green-800 focus-visible:ring-green-600/30 shadow-sm font-semibold dark:bg-green-500 dark:hover:bg-green-600",
        "outline-success": 
          "border-2 border-green-600 text-green-600 bg-transparent hover:bg-green-50 active:bg-green-100 focus-visible:ring-green-600/30 font-semibold dark:border-green-400 dark:text-green-400 dark:hover:bg-green-950",
      },
      size: {
        default: "h-9 px-4 py-2 has-[>svg]:px-3",
        sm: "h-8 rounded-md gap-1.5 px-3 has-[>svg]:px-2.5",
        lg: "h-10 rounded-md px-6 has-[>svg]:px-4",
        xl: "h-12 rounded-lg px-8 has-[>svg]:px-6 text-base", // Added for prominent buttons
        icon: "size-9",
      },
    },
    defaultVariants: {
      variant: "default",
      size: "default",
    },
  }
)

function Button({
  className,
  variant,
  size,
  asChild = false,
  ...props
}: React.ComponentProps<"button"> &
  VariantProps<typeof buttonVariants> & {
    asChild?: boolean
  }) {
  const Comp = asChild ? Slot : "button"

  return (
    <Comp
      data-slot="button"
      className={cn(buttonVariants({ variant, size, className }))}
      {...props}
    />
  )
}

export { Button, buttonVariants }