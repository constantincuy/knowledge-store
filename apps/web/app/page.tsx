import {Metadata} from "next"

import {MainNav} from "@/app/components/main-nav"
import {UserNav} from "@/app/components/user-nav";
import {SearchArea} from "@/app/components/search-area";

export const metadata: Metadata = {
    title: "Document Search",
}


export default function DashboardPage() {
    return (
        <div className="flex-col">
            <div className="border-b">
                <div className="flex h-16 items-center px-4">
                    <MainNav className="mx-6"/>
                    <div className="ml-auto flex items-center space-x-4">
                        <UserNav/>
                    </div>
                </div>
            </div>
            <SearchArea />
        </div>
    )
}