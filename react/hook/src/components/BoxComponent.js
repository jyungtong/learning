import { useTheme } from '../contexts/ThemeContext';
import './BoxComponent.css';

export default function BoxComponent() {
  const isRedTheme = useTheme();
  const style = isRedTheme ? 'redBox' : 'greyBox';

  console.log('====BoxComponent rerender');

  return (
    <div className={`${style} box`}>
      A Box
    </div>
  );
}