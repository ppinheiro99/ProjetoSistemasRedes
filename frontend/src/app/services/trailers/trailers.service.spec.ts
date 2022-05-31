import { TestBed } from '@angular/core/testing';

import { TrailersService } from './trailers.service';

describe('TrailerService', () => {
  let service: TrailersService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(TrailersService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
