CREATE TABLE videogames(
	id SERIAL PRIMARY KEY,
	name VARCHAR(100),
	category VARCHAR(100),
	active_players INT,
	size FLOAT,
	rating INT,
	downloads INT
);

INSERT INTO videogames (id, name, category, active_players, size, rating, downloads) VALUES
(1, 'Fortnite', 'Battle Royale', 5000000, 30.5, 9, 400000000),
(2, 'Minecraft', 'Sandbox', 3000000, 1.2, 10, 238000000),
(3, 'Call of Duty Warzone', 'Shooter', 2000000, 80.0, 8, 100000000),
(4, 'League of Legends', 'MOBA', 2500000, 15.0, 9, 180000000),
(5, 'Valorant', 'Shooter', 1500000, 25.0, 9, 50000000),
(6, 'GTA V', 'Action', 1200000, 95.0, 10, 170000000),
(7, 'Roblox', 'Sandbox', 4000000, 0.3, 8, 200000000),
(8, 'Apex Legends', 'Battle Royale', 1100000, 60.0, 8, 100000000),
(9, 'Among Us', 'Party', 500000, 0.2, 7, 500000000),
(10, 'PUBG', 'Battle Royale', 900000, 40.0, 8, 75000000),
(11, 'Clash of Clans', 'Strategy', 800000, 0.5, 9, 600000000),
(12, 'Clash Royale', 'Strategy', 700000, 0.4, 8, 500000000),
(13, 'The Witcher 3', 'RPG', 300000, 50.0, 10, 50000000),
(14, 'Cyberpunk 2077', 'RPG', 350000, 70.0, 7, 30000000),
(15, 'Elden Ring', 'RPG', 600000, 60.0, 10, 25000000),
(16, 'FIFA 23', 'Sports', 900000, 45.0, 8, 30000000),
(17, 'NBA 2K23', 'Sports', 400000, 55.0, 7, 20000000),
(18, 'Animal Crossing', 'Simulation', 200000, 6.5, 9, 35000000),
(19, 'Stardew Valley', 'Simulation', 150000, 0.5, 9, 20000000),
(20, 'Terraria', 'Adventure', 180000, 0.4, 9, 44000000),
(21, 'Halo Infinite', 'Shooter', 250000, 75.0, 8, 20000000),
(22, 'Overwatch 2', 'Shooter', 500000, 50.0, 7, 35000000),
(23, 'Genshin Impact', 'RPG', 1000000, 35.0, 9, 60000000),
(24, 'Rocket League', 'Sports', 300000, 20.0, 8, 75000000),
(25, 'Fall Guys', 'Party', 220000, 15.0, 8, 50000000);

SELECT setval('videogames_id_seq', (SELECT MAX(id) FROM videogames));