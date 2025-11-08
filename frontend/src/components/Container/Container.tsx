import cn from 'classnames';
import type { CSSProperties, ReactNode } from 'react';
import styles from './Container.module.css';

type ContainerProps = {
  children: ReactNode;
  className?: string;
  style?: CSSProperties;
};

const Container = ({ children, className, style }: ContainerProps) => {
  return (
    <div className={cn(styles.container, className)} style={style}>
      {children}
    </div>
  );
};

export default Container;
