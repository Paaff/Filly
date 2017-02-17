import { FillyPage } from './app.po';

describe('filly App', function() {
  let page: FillyPage;

  beforeEach(() => {
    page = new FillyPage();
  });

  it('should display message saying app works', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('app works!');
  });
});
