export default function generateName() {
    const name = names[Math.floor(Math.random() * names.length)];
    const adjective = adjectives[Math.floor(Math.random() * adjectives.length)];

    return `${adjective} ${name}`
}


const names = [
    "Apple",
    "Ash",
    "Basil",
    "Bear",
    "Birch",
    "Blaze",
    "Book",
    "Breeze",
    "Brook",
    "Cedar",
    "Clay",
    "Cliff",
    "Cloud",
    "Coral",
    "Cove",
    "Crow",
    "Dawn",
    "Deer",
    "Dell",
    "Drake",
    "Dune",
    "Earth",
    "East",
    "Echo",
    "Ember",
    "Falcon",
    "Fern",
    "Fern",
    "Finch",
    "Fire",
    "Fjord",
    "Flame",
    "Flint",
    "Fox",
    "Frost",
    "Glade",
    "Glen",
    "Grove",
    "Hare",
    "Hawk",
    "Hazel",
    "Hollow",
    "Ice",
    "Iris",
    "Ivy",
    "Jade",
    "Jet",
    "Juniper",
    "Kite",
    "Lavender",
    "Leaf",
    "Light",
    "Lilac",
    "Lily",
    "Loop",
    "Lynx",
    "Maple",
    "Meadow",
    "Mint",
    "Mist",
    "Moon",
    "Moss",
    "Oak",
    "Ocean",
    "Olive",
    "Onyx",
    "Opal",
    "Pearl",
    "Pine",
    "Pool",
    "Quartz",
    "Rain",
    "Raven",
    "Ridge",
    "River",
    "Robin",
    "Rocket",
    "Rose",
    "Ruby",
    "Sable",
    "Sage",
    "Sand",
    "Shadow",
    "Sky",
    "Slate",
    "Snow",
    "Star",
    "Stone",
    "Storm",
    "Sun",
    "Swan",
    "Thunder",
    "Thyme",
    "Vale",
    "Violet",
    "Wave",
    "Willow",
    "Wind",
    "Wolf",
    "Wren",
]

const adjectives = ["Adventurous",
    "Ambitious",
    "Brilliant",
    "Bold",
    "Bright",
    "Cheerful",
    "Confident",
    "Creative",
    "Curious",
    "Daring",
    "Delightful",
    "Determined",
    "Dynamic",
    "Energetic",
    "Enthusiastic",
    "Exuberant",
    "Fearless",
    "Friendly",
    "Fun-loving",
    "Genuine",
    "Gifted",
    "Happy",
    "Harmonious",
    "Heroic",
    "Honest",
    "Hopeful",
    "Humble",
    "Imaginative",
    "Innovative",
    "Inspiring",
    "Intelligent",
    "Intrepid",
    "Joyful",
    "Kind",
    "Lively",
    "Loyal",
    "Magnanimous",
    "Magical",
    "Majestic",
    "Marvelous",
    "Merry",
    "Mindful",
    "Motivated",
    "Nimble",
    "Optimistic",
    "Outgoing",
    "Passionate",
    "Peaceful",
    "Perceptive",
    "Persistent",
    "Playful",
    "Positive",
    "Powerful",
    "Precise",
    "Proactive",
    "Progressive",
    "Radiant",
    "Resilient",
    "Resourceful",
    "Responsible",
    "Revolutionary",
    "Robust",
    "Romantic",
    "Sincere",
    "Skilled",
    "Sophisticated",
    "Spirited",
    "Strong",
    "Talented",
    "Tenacious",
    "Thoughtful",
    "Tireless",
    "Tranquil",
    "Trustworthy",
    "Unique",
    "Unwavering",
    "Vibrant",
    "Vigilant",
    "Virtuous",
    "Vivacious",
    "Warm",
    "Wise",
    "Witty",
    "Wonderful",
    "Xenial",
    "Youthful",
    "Zealous",
    "Zestful",
    "Boldly",
    "Brilliantly",
    "Cheerfully",
    "Confidently",
    "Creatively",
    "Dazzlingly",
    "Dynamically",
    "Energetically",
    "Enthusiastically",
    "Fearlessly",
    "Flamboyantly",
    "Radiantly"]
