import { ComponentFixture, TestBed } from '@angular/core/testing';

import { PassrecoverFormComponent } from './passrecover-form.component';

describe('PassrecoverFormComponent', () => {
  let component: PassrecoverFormComponent;
  let fixture: ComponentFixture<PassrecoverFormComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ PassrecoverFormComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PassrecoverFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
