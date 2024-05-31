import {
    benefitIcon1,
    benefitIcon2,
    benefitIcon3,
    benefitIcon4,
    benefitImage2,
    chromecast,
    disc02,
    discord,
    discordBlack,
    facebook,
    figma,
    file02,
    framer,
    homeSmile,
    instagram,
    notification2,
    notification3,
    notification4,
    notion,
    photoshop,
    plusSquare,
    protopie,
    raindrop,
    recording01,
    recording03,
    roadmap1,
    roadmap2,
    roadmap3,
    roadmap4,
    searchMd,
    slack,
    sliders04,
    telegram,
    twitter,
    yourlogo,
  } from "../assets";
  
  export const navigation = [
    {
      id: "0",
      title: "Features",
      url: "/#features",
    },
    {
      id: "1",
      title: "Pricing",
      url: "/#pricing",
    },
   
    {
      id: "4",
      title: "New account",
      url: "/register",
      onlyMobile: true,
    },
    {
      id: "5",
      title: "Sign in",
      url: "/login",
      onlyMobile: true,
    },
  ];

  export const loggedInNavigation = [
    {
        id: "0",
        title: "Home",
        url: "/mainmenu",
        },

    {
        id: "1",
        title: "My Courses",
        url: "/mycourses",
        },

    {
        id: "2",
        title: "My Profile",
        url: "/profile",
        },
    
  ];

  export const loggedInActions = [
    { id: '1', url: '/dashboard', title: 'Dashboard' },
    { id: '2', url: '/profile', title: 'Profile' },
    { id: '3', url: '/settings', title: 'Settings' },
  ];
  
  export const heroIcons = [homeSmile, file02, searchMd, plusSquare];
  
  export const notificationImages = [notification4, notification3, notification2];
  
  export const companyLogos = [yourlogo, yourlogo, yourlogo, yourlogo, yourlogo];
  
  export const brainwaveServices = [
    "Photo generating",
    "Photo enhance",
    "Seamless Integration",
  ];
  
  export const brainwaveServicesIcons = [
    recording03,
    recording01,
    disc02,
    chromecast,
    sliders04,
  ];
  
 
  
 
 
  
 
  export const pricing = [
    {
      id: "0",
      title: "Basic",
      description: "Free courses, but no teacher support :C",
      price: "0.00",
      features: [
        "Great courses to learn from",
        "Get certificates for completing courses",
        "Ability to explore the app and its features without any cost",
      ],
    },
    {
      id: "1",
      title: "Premium",
      description: "All of the magic of FREMICOurses, priority support, and more!",
      price: "9.99 / course",
      features: [
        "Access to all courses on the platform",
        "Get certificates for completing courses",
        "Priority support to solve issues quickly",
      ],
    },
    {
      id: "2",
      title: "Enterprise",
      description: "Dedicated courses",
      price: null,
      features: [
        "Fully tailored courses for your organization",
        "Personalized recommendations based on your preferences",
        "Ability to explore the app and its features without any cost",
      ],
    },
  ];
  
  export const benefits = [
    {
      id: "0",
      title: "Learn anything",
      text: "Lets users quickly find answers to their questions with our elite courses.",
      backgroundUrl: "./src/assets/benefits/card-1.svg",
      iconUrl: benefitIcon1,
      imageUrl: benefitImage2,
    },
    {
      id: "1",
      title: "Improve everyday",
      text: "Our courses are updated everyday to keep you up to date with the latest technologies.",
      backgroundUrl: "./src/assets/benefits/card-2.svg",
      iconUrl: benefitIcon2,
      imageUrl: benefitImage2,
      light: true,
    },
   
    {
      id: "2",
      title: "Benefit from interactions with professionals",
      text: "Our courses are taught by professionals in the field who have experience in the topics they teach.",
      backgroundUrl: "./src/assets/benefits/card-4.svg",
      iconUrl: benefitIcon4,
      imageUrl: benefitImage2,
      light: true,
    },
   
  ];
  
  export const socials = [
    {
      id: "0",
      title: "Discord",
      iconUrl: discordBlack,
      url: "#",
    },
    {
      id: "1",
      title: "Twitter",
      iconUrl: twitter,
      url: "#",
    },
    {
      id: "2",
      title: "Instagram",
      iconUrl: instagram,
      url: "#",
    },
    {
      id: "3",
      title: "Telegram",
      iconUrl: telegram,
      url: "#",
    },
    {
      id: "4",
      title: "Facebook",
      iconUrl: facebook,
      url: "#",
    },
  ];