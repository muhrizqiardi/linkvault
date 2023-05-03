import { afterAll, afterEach, beforeAll, describe, it } from '@jest/globals';
import { fireEvent, render, screen } from '@testing-library/react';
import { RegisterPageForm } from './register-page-form';
import { faker } from '@faker-js/faker';
import { setupServer } from 'msw/node';
import { rest } from 'msw';

const server = setupServer(
  rest.post('/api/users', (_res, res, ctx) =>
    res(
      ctx.json({
        success: true,
        message: 'User registered',
        data: {
          created_at: new Date().toISOString(),
          updated_at: new Date().toISOString(),
          full_name: faker.name.fullName(),
          email: faker.internet.email(),
          id: faker.datatype.uuid(),
        },
      }),
    ),
  ),
);

describe('RegisterPageForm', () => {
  beforeAll(() => server.listen());
  afterEach(() => server.resetHandlers());
  afterAll(() => server.close());

  it('should display error message when form are submitted but password is less than 8 chars', async () => {
    render(<RegisterPageForm />);

    const fullNameInputElement = screen.getByLabelText(/full name/i);
    const emailInputElement = screen.getByLabelText(/email/i);
    const passwordInputElement = screen.getByLabelText('Password');
    const confirmPasswordInputElement =
      screen.getByLabelText(/confirm password/i);
    const createAccountButtonElement = screen.getByText(/create account/i);

    fireEvent.change(fullNameInputElement, {
      target: {
        value: faker.name.fullName(),
      },
    });
    fireEvent.change(emailInputElement, {
      target: {
        value: faker.internet.email(),
      },
    });
    fireEvent.change(passwordInputElement, {
      target: {
        value: '5char',
      },
    });
    fireEvent.change(confirmPasswordInputElement, {
      target: {
        value: {
          value: '5char',
        },
      },
    });
    fireEvent.click(createAccountButtonElement);

    await screen.findByText('Password should have more than 8 characters');
  });

  it("should display error message when password and confirm password doesn't match", async () => {
    render(<RegisterPageForm />);

    const fullNameInputElement = screen.getByLabelText(/full name/i);
    const emailInputElement = screen.getByLabelText(/email/i);
    const passwordInputElement = screen.getByLabelText('Password');
    const confirmPasswordInputElement =
      screen.getByLabelText(/confirm password/i);
    const createAccountButtonElement = screen.getByText(/create account/i);

    fireEvent.change(fullNameInputElement, {
      target: {
        value: faker.name.fullName(),
      },
    });
    fireEvent.change(emailInputElement, {
      target: {
        value: faker.internet.email(),
      },
    });
    fireEvent.change(passwordInputElement, {
      target: {
        value: '5char',
      },
    });
    fireEvent.change(confirmPasswordInputElement, {
      target: {
        value: {
          value: '5char',
        },
      },
    });
    fireEvent.click(createAccountButtonElement);

    await screen.findByText("Password confirmation doesn't match");
  });
});
