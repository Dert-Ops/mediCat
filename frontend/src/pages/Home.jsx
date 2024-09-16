import { useState, useEffect, useCallback } from 'react';
import Navbar from '../components/Navbar';

export default function Home() {
	const [posts, setPosts] = useState([
		{
			id: 1,
			user: 'john_doe',
			imageUrl: 'https://via.placeholder.com/400',
			caption: 'Beautiful sunset at the beach!',
			comments: [
				{ user: 'jane_doe', text: 'Amazing view!' },
				{ user: 'mark_smith', text: 'Wish I was there!' }
			]
		},
		// Başlangıçta birkaç tane daha post ekleyin
		{
			id: 2,
			user: 'alice_wonder',
			imageUrl: 'https://via.placeholder.com/400',
			caption: 'Hiking adventure in the mountains.',
			comments: [
				{ user: 'bob_brown', text: 'Looks like fun!' },
				{ user: 'lucy_liu', text: 'Great shot!' }
			]
		},
		{
			id: 3,
			user: 'emma_jones',
			imageUrl: 'https://via.placeholder.com/400',
			caption: 'Delicious homemade pizza!',
			comments: [
				{ user: 'david_clark', text: 'Yummy!' },
				{ user: 'susan_davis', text: 'Looks delicious!' }
			]
		}
	]);
	const [loading, setLoading] = useState(false);

	const fetchMorePosts = useCallback(() => {
		if (loading) return;
		setLoading(true);

		// Yeni postlar ekleyin (backend'den geliyormuş gibi)
		const newPosts = [
			{
				id: 1,
				user: 'john_doe',
				imageUrl: 'https://via.placeholder.com/400',
				caption: 'Beautiful sunset at the beach!',
				comments: [
					{ user: 'jane_doe', text: 'Amazing view!' },
					{ user: 'mark_smith', text: 'Wish I was there!' }
				]
			},
			{
				id: 2,
				user: 'alice_wonder',
				imageUrl: 'https://via.placeholder.com/400',
				caption: 'Hiking adventure in the mountains.',
				comments: [
					{ user: 'bob_brown', text: 'Looks like fun!' },
					{ user: 'lucy_liu', text: 'Great shot!' }
				]
			},
			{
				id: 3,
				user: 'john_doe',
				imageUrl: 'https://via.placeholder.com/400',
				caption: 'Beautiful sunset at the beach!',
				comments: [
					{ user: 'jane_doe', text: 'Amazing view!' },
					{ user: 'mark_smith', text: 'Wish I was there!' }
				]
			},
			{
				id: 4,
				user: 'alice_wonder',
				imageUrl: 'https://via.placeholder.com/400',
				caption: 'Hiking adventure in the mountains.',
				comments: [
					{ user: 'bob_brown', text: 'Looks like fun!' },
					{ user: 'lucy_liu', text: 'Great shot!' }
				]
			},
			{
				id: 5,
				user: 'john_doe',
				imageUrl: 'https://via.placeholder.com/400',
				caption: 'Beautiful sunset at the beach!',
				comments: [
					{ user: 'jane_doe', text: 'Amazing view!' },
					{ user: 'mark_smith', text: 'Wish I was there!' }
				]
			},
			{
				id: 6,
				user: 'alice_wonder',
				imageUrl: 'https://via.placeholder.com/400',
				caption: 'Hiking adventure in the mountains.',
				comments: [
					{ user: 'bob_brown', text: 'Looks like fun!' },
					{ user: 'lucy_liu', text: 'Great shot!' }
				]
			},
			{
				id: 7,
				user: 'john_doe',
				imageUrl: 'https://via.placeholder.com/400',
				caption: 'Beautiful sunset at the beach!',
				comments: [
					{ user: 'jane_doe', text: 'Amazing view!' },
					{ user: 'mark_smith', text: 'Wish I was there!' }
				]
			},
			{
				id: 8,
				user: 'alice_wonder',
				imageUrl: 'https://via.placeholder.com/400',
				caption: 'Hiking adventure in the mountains.',
				comments: [
					{ user: 'bob_brown', text: 'Looks like fun!' },
					{ user: 'lucy_liu', text: 'Great shot!' }
				]
			},
			{
				id: 9,
				user: 'john_doe',
				imageUrl: 'https://via.placeholder.com/400',
				caption: 'Beautiful sunset at the beach!',
				comments: [
					{ user: 'jane_doe', text: 'Amazing view!' },
					{ user: 'mark_smith', text: 'Wish I was there!' }
				]
			},
			{
				id: 10,
				user: 'alice_wonder',
				imageUrl: 'https://via.placeholder.com/400',
				caption: 'Hiking adventure in the mountains.',
				comments: [
					{ user: 'bob_brown', text: 'Looks like fun!' },
					{ user: 'lucy_liu', text: 'Great shot!' }
				]
			},
			{
				id: 11,
				user: 'john_doe',
				imageUrl: 'https://via.placeholder.com/400',
				caption: 'Beautiful sunset at the beach!',
				comments: [
					{ user: 'jane_doe', text: 'Amazing view!' },
					{ user: 'mark_smith', text: 'Wish I was there!' }
				]
			},
			{
				id: 12,
				user: 'alice_wonder',
				imageUrl: 'https://via.placeholder.com/400',
				caption: 'Hiking adventure in the mountains.',
				comments: [
					{ user: 'bob_brown', text: 'Looks like fun!' },
					{ user: 'lucy_liu', text: 'Great shot!' }
				]
			},
		];

		setTimeout(() => {
			setPosts((prevPosts) => [...prevPosts, ...newPosts]);
			setLoading(false);
		}, 1000); // Yeni postları 1 saniye gecikme ile ekleyin (yükleniyormuş gibi)
	}, [posts, loading]);

	const handleScroll = () => {
		if (window.innerHeight + document.documentElement.scrollTop >= document.documentElement.offsetHeight - 200) {
			fetchMorePosts();
		}
	};

	useEffect(() => {
		window.addEventListener('scroll', handleScroll);
		return () => window.removeEventListener('scroll', handleScroll);
	}, [handleScroll]);

	return (
		<>
			<Navbar />
			return (
			<div className="container mx-auto px-4">
				<div className="flex flex-col items-center gap-4">
					{posts.map((post) => (
						<div key={post.id} className="bg-white p-4 rounded shadow w-[800px] max-w-full">
							<div className="flex items-center mb-4">
								<div className="bg-gray-200 rounded-full w-10 h-10 flex items-center justify-center">
									<span className="text-gray-800 font-bold">{post.user[0].toUpperCase()}</span>
								</div>
								<p className="ml-2 text-gray-700 font-bold">{post.user}</p>
							</div>
							<img src={post.imageUrl} alt={post.caption} className="w-full h-auto rounded mb-2" />
							<p className="text-gray-700 mb-2">{post.caption}</p>
							<div>
								<h4 className="text-gray-600 font-bold">Comments:</h4>
								{post.comments.map((comment, index) => (
									<p key={index} className="text-gray-600">
										<span className="font-bold">{comment.user}:</span> {comment.text}
									</p>
								))}
							</div>
						</div>
					))}
				</div>
				{loading && (
					<div className="flex justify-center items-center h-16">
						<div className="flex items-center space-x-2">
							<div className="w-6 h-6 border-4 border-blue-500 border-t-transparent border-solid rounded-full animate-spin"></div>
							<p className="text-blue-500 font-semibold text-lg">Loading...</p>
						</div>
					</div>
				)}
			</div>
		</>
	);
}
