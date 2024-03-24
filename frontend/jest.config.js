module.exports = {
    testEnvironment: "jsdom",
    transform: {
        '^.+\\.ts$': 'babel-jest',
		'^.+\\.js$': 'babel-jest',
		'^.+\\.vue$': '@vue/vue3-jest',
    },
    testRegex: "(/__tests__/.*|(\\.|/)(test|spec))\\.(js|ts)$",
    moduleFileExtensions: ['js', 'json', 'vue', 'ts'],
    moduleNameMapper: {
        "^@/(.*)$": "<rootDir>/src/$1",
    },
    coveragePathIgnorePatterns: ["/node_modules/", "/tests/"],
    coverageReporters: ["text", "json-summary"],
    setupFiles: ['<rootDir>/jest-setup.ts'], 
    testEnvironmentOptions: {
        customExportConditions: ["node"],
    },
}