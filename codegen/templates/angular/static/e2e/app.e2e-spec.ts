import { GoRamlPage } from './app.po';

describe('go-raml App', () => {
  let page: GoRamlPage;

  beforeEach(() => {
    page = new GoRamlPage();
  });

  it('should display message saying app works', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('app works!');
  });
});
