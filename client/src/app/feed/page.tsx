"use client";

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import type { Post, FeedResponse } from '@/interface/PostInterface';
import { useAuth } from '@/context/AuthContext';

const Feed = () => {
  const [posts, setPosts] = useState<Post[]>([]);
  const [editingPostId, setEditingPostId] = useState<number | null>(null);
  const [editedContent, setEditedContent] = useState<string>('');
  const [newPostContent, setNewPostContent] = useState<string>('');
  const { isAuthenticated } = useAuth();
  const router = useRouter();

  const fetchPosts = async () => {
    try {
      const response = await fetch('http://localhost:8080/api/post/');
      const data: FeedResponse = await response.json();
      setPosts(data.posts);
    } catch (error) {
      console.error('Failed to fetch posts', error);
    }
  };

  const handleCreatePost = async () => {
    try {
      const response = await fetch('http://localhost:8080/api/post/create', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${localStorage.getItem('token-ssm')}`,
        },
        body: JSON.stringify({ content: newPostContent }),
      });

      if (response.ok) {
        setNewPostContent('');
        fetchPosts();
      } else {
        alert('Failed to create post');
      }
    } catch (error) {
      alert('Failed to create post');
    }
  };

  const handleEdit = (postId: number, currentContent: string) => {
    setEditingPostId(postId);
    setEditedContent(currentContent);
  };

  const handleUpdate = async (postId: number) => {
    try {
      const response = await fetch(`http://localhost:8080/api/post/${postId}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${localStorage.getItem('token-ssm')}`,
        },
        body: JSON.stringify({ content: editedContent }),
      });

      if (response.ok) {
        setEditingPostId(null);
        fetchPosts();
      } else {
        alert('Failed to update post');
      }
    } catch (error) {
      alert('Failed to update post');
    }
  };

  const handleDelete = async (postId: number) => {
    try {
      const response = await fetch(`http://localhost:8080/api/post/${postId}`, {
        method: 'DELETE',
        headers: {
          Authorization: `Bearer ${localStorage.getItem('token-ssm')}`,
        },
      });

      if (response.ok) {
        fetchPosts();
      } else {
        alert('Failed to delete post');
      }
    } catch (error) {
      alert('Failed to delete post');
    }
  };

  useEffect(() => {
    fetchPosts();
  }, [router]);

  return (
    <div className="flex min-h-screen bg-gray-100">
      <div className="max-w-4xl mx-auto p-4">
        <h1 className="text-3xl font-bold mb-6 text-black text-center">Threads</h1>
        {
          (!isAuthenticated) ? (
            <div className="bg-white mb-6 p-4 rounded-lg shadow-md">
              <p className="text-center text-black">Please login to create a new post</p>
            </div>
          ) : (
            <div className="mb-6 p-4 bg-white rounded-lg shadow-md">
              <textarea
                className="w-full p-2 border border-gray-300 rounded-lg text-black mb-2"
                placeholder="Write a new post..."
                value={newPostContent}
                onChange={(e) => setNewPostContent(e.target.value)}
              />
              <button
                onClick={handleCreatePost}
                className="bg-blue-600 text-white py-2 px-4 rounded-lg hover:bg-blue-800"
              >
                Create Post
              </button>
            </div>
          )
        }
        
        <div className="space-y-4">
          {posts.map(post => (
            <div key={post.id} className="p-4 bg-white rounded-lg shadow-md">
              <div className="flex justify-between items-center mb-2">
                <div>
                  <span className="font-semibold text-black">{post.user.displayName}</span> 
                  <span className="text-gray-500 text-sm ml-2">
                    (@{post.user.username}) {new Date(post.createdAt).toLocaleDateString()}
                  </span>
                </div>

                <div className="flex space-x-2">
                  {editingPostId === post.id ? (
                    <>
                      <button
                        onClick={() => handleUpdate(post.id)}
                        className="text-blue-600 hover:text-blue-800"
                      >
                        Update Post
                      </button>
                      <button
                        onClick={() => setEditingPostId(null)}
                        className="text-red-600 hover:text-red-800"
                      >
                        Cancel
                      </button>
                    </>
                  ) : (
                    <>
                      <button
                        onClick={() => handleEdit(post.id, post.content)}
                        className="text-blue-600 hover:text-blue-800"
                      >
                        Edit
                      </button>
                      <button
                        onClick={() => handleDelete(post.id)}
                        className="text-red-600 hover:text-red-800"
                      >
                        Delete
                      </button>
                    </>
                  )}
                </div>
                
              </div>
              {editingPostId === post.id ? (
                <textarea
                  className="w-full p-2 border border-gray-300 rounded-lg text-black"
                  value={editedContent}
                  onChange={(e) => setEditedContent(e.target.value)}
                />
              ) : (
                <p className="text-black break-words">{post.content}</p>
              )}
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};

export default Feed;
