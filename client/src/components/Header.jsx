import { useState } from "react";
import { FaMoon, FaSun } from "react-icons/fa";
import { motion } from "framer-motion";

const Header = () => {
  const [isDarkMode, setIsDarkMode] = useState(false);

  const handleThemeToggle = () => {
    setIsDarkMode(!isDarkMode);
    if (!isDarkMode) {
      document.documentElement.classList.add("dark");
    } else {
      document.documentElement.classList.remove("dark");
    }
  };

  return (
    <header className="w-full p-4 bg-primary text-white shadow-lg">
      <div className="container mx-auto flex items-center justify-between">
        {/* Left Side: App Title */}
        <motion.div
          className="text-2xl font-bold"
          initial={{ x: -100, opacity: 0 }}
          animate={{ x: 0, opacity: 1 }}
          transition={{ duration: 0.5 }}
        >
          TodoApp
        </motion.div>

        {/* Center: Tagline */}
        <motion.div
          className="text-lg italic"
          initial={{ opacity: 0 }}
          animate={{ opacity: 1 }}
          transition={{ duration: 0.5 }}
        >
          Get things done, one task at a time.
        </motion.div>

        {/* Right Side: Theme Switcher */}
        <motion.div
          className="flex items-center space-x-4"
          initial={{ opacity: 0 }}
          animate={{ opacity: 1 }}
          transition={{ duration: 0.5 }}
        >
          <button
            onClick={handleThemeToggle}
            className="text-2xl focus:outline-none"
          >
            {isDarkMode ? (
              <FaSun className="text-yellow-400" />
            ) : (
              <FaMoon className="text-blue-400" />
            )}
          </button>
        </motion.div>
      </div>
    </header>
  );
};

export default Header;
