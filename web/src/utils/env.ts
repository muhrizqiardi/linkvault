import { str, envsafe } from 'envsafe';

export const env = envsafe({
  API_URL: str({
    devDefault: 'http://localhost:9000',
  }),
});
