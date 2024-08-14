import type { User } from "./UserInterface";

interface Post {
    id: number;
    content: string;
    createdAt: string;
    updatedAt: string;
    user: User;
}

interface FeedResponse {
    posts: Post[];
}

export type { Post, FeedResponse };