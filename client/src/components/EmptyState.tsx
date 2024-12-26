
import React from "react";

const EmptyState = () => (
  <div className="flex flex-col items-center justify-center p-8 text-center">
    <div className="w-16 h-16 bg-gray-200 dark:bg-gray-700 rounded-full flex items-center justify-center mb-4">
      <span className="text-2xl">🏏</span>
    </div>
    <h3 className="text-xl font-semibold text-gray-800 dark:text-gray-200 mb-2">
      No Videos Found
    </h3>
    <p className="text-gray-600 dark:text-gray-400">
      Check back later for new cricket highlights!
    </p>
  </div>
);

export default EmptyState;
