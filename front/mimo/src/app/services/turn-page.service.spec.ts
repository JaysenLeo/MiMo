import { TestBed } from '@angular/core/testing';

import { TurnPageService } from './turn-page.service';

describe('TurnPageService', () => {
  let service: TurnPageService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(TurnPageService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
