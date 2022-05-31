import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PassrecoverComponent } from './passrecover.component';

describe('PassrecoverComponent', () => {
  let component: PassrecoverComponent;
  let fixture: ComponentFixture<PassrecoverComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PassrecoverComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PassrecoverComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
