import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import MealPlan from './components/MealPlan';
import Tables from './components/Table';

export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/welcome', WelcomePage);
    router.registerRoute('/meal_plan', MealPlan);
    router.registerRoute('/', Tables);
  },
});
