"use client"
import { Input } from "@/components/ui/input"
import {InputHTMLAttributes} from "react";

interface SearchProps extends InputHTMLAttributes<HTMLInputElement>{
}

export function Search(props: SearchProps) {
    return (
        <Input
            type="search"
            placeholder="e.g. Sales documents"
            {...props}
        />
    )
}